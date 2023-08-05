#ifndef _richdem_topological_sorts_
#define _richdem_topological_sorts_

#include <queue>
#include <richdem/common/Array3D.hpp>
#include <richdem/common/logger.hpp>
#include <richdem/common/ProgressBar.hpp>
#include <vector>

namespace richdem {


///The donors of a focal cell are the neighbours from which it receives flow.
///Here, we identify those neighbours by inverting the Receivers array.
void ComputeDonors(){
  //Initially, we claim that each cell has no donors.
  for(int i=0;i<size;i++)
    ndon[i] = 0;

  //Looping across all cells
  for(int c=0;c<size;c++){
    if(rec[c]==NO_FLOW)
      continue;
    //If this cell passes flow to a downhill cell, make a note of it in that
    //downhill cell's donor array and increment its donor counter
    const auto n       = c+nshift[rec[c]];
    donor[8*n+ndon[n]] = c;
    ndon[n]++;
  }
}









static void FindStack(std::vector<int> &ordering, const int c, int &nstack){
  for(int k=0;k<ndon[c];k++){
    const auto n    = donor[8*c+k];
    stack[nstack++] = n;
    FindStack(n,nstack);
  }
}




/**
  @brief  Calculate a depth-first traversal topology
  @author Richard Barnes (rbarnes@umn.edu)

  Given a flow proportions array, calculate a depth-first traversal of it. Only
  valid for non-divergent flow metrics.
   
  @param[in]     props        Flow proportions array
  @param[out]    &ordering    Output vector for the cell ordering
  @param[out]    &levels      Divisions of the output vector into separably processible units

  @pre
    1. The flow proportions array contains information from a 
       non-divergent flow metric.

    2. The flow proportions matrix must already be initialized.

  @post
    1. \p ordering is modified to contain the indices of cells ordered such that
       a traversal from either end of the vector to the other propagates
       information correctly.

    2. levels will contain indices indicating elements of \p ordering. Adjacent
       indices form a half-open interval `[start,end)` whose cells can be
       processed in parallel with other half-open intervals. The last value of 
       the levels array is constructed so as to form a valid interval.
*/
template<class A>
void TopologicalSortDepthFirst(
  const Array3D<float> &props,
  std::vector<int>     &ordering,
  std::vector<int>     &levels
){
  Timer overall;
  overall.start();

  RDLOG_ALG_NAME<<"Depth-First Topological Sort";

  //The `levels` array has an unpredictable size, so we fill it dynamically and
  //reset it each time. This should have minimal impact on the algorithm's speed
  //since std::vector's memory is not actually reallocated.
  stack_start.clear();

  int nstack=0;
  for(auto c=props.i0();c<props.size();c++){
    if(props.getIN(c,0)==NO_FLOW_GEN){
      stack_start.push_back(nstack);
      ordering[nstack++] = c;
      FindStack(ordering,c,nstack);
    }
  }  
  //We add an additional note to the end of `stack_start` that serves as an
  //upper bound on the locations of the cells in the final ordering. See b    
  stack_start.push_back(nstack);

}
















  

  //Create dependencies array
  RDLOG_PROGRESS<<"Creating dependencies array..."<<std::endl;
  Array2D<int8_t> deps(props, 0);
  for(int y=1;y<props.height()-1;y++)
  for(int x=1;x<props.width()-1;x++){
    const int ci = accum.xyToI(x,y);
    if(props.isNoData(ci))
      continue;
    for(int n=1;n<=8;n++)
      if(props(x,y,n)>0){
        const int ni = ci + accum.nshift(n);
        deps(ni)++;
      }
  }

  //Find sources
  std::queue<int> q;
  for(auto i=deps.i0();i<deps.size();i++)
    if(deps(i)==0 && !props.isNoData(i))
      q.emplace(i);

  RDLOG_DEBUG<<"Source cells found = "<<q.size(); //TODO: Switch log target

  RDLOG_PROGRESS<<"Calculating flow accumulation...";
  ProgressBar progress;
  progress.start(props.size());
  while(!q.empty()){
    ++progress;

    const auto ci = q.front();
    q.pop();

    assert(!props.isNoData(ci));

    const auto c_accum = accum(ci);

    for(int n=1;n<=8;n++){
      if(props.getIN(ci,n)<=0) //No Flow in this direction or other flags
        continue;
      const int ni = ci+accum.nshift(n);
      if(props.isNoData(ni))
        continue;
      accum(ni) += props.getIN(ci,n)*c_accum;
      if(--deps(ni)==0)
        q.emplace(ni);
      assert(deps(ni)>=0);
    }
  }
  progress.stop();

  for(auto i=props.i0();i<props.size();i++)
    if(props.isNoData(i))
      accum(i) = accum.noData();

  RDLOG_TIME_USE<<"Wall-time       = "<<overall.stop()<<" s"     ;
}












  ///Cells must be ordered so that they can be traversed such that higher cells
  ///are processed before their lower neighbouring cells. This method creates
  ///such an order. It also produces a list of "levels": cells which are,
  ///topologically, neither higher nor lower than each other. Cells in the same
  ///level can all be processed simultaneously without having to worry about
  ///race conditions.
  void GenerateOrder(){
    int nstack = 0;    //Number of cells currently in the stack

    //Since each value of the `levels` array is later used as the starting value
    //of a for-loop, we include a zero at the beginning of the array.
    levels[0] = 0;     
    nlevel    = 1;     //Note that array now contains a single value

    //Load cells without dependencies into the queue. This will include all of
    //the edge cells.
    for(int c=0;c<size;c++){
      if(rec[c]==NO_FLOW)
        stack[nstack++] = c;
    }
    levels[nlevel++] = nstack; //Last cell of this level

    //Start with level_bottom=-1 so we get into the loop, it is immediately
    //replaced by level_top.
    int level_bottom = -1;         //First cell of the current level
    int level_top    =  0;         //Last cell of the current level

    while(level_bottom<level_top){ //Enusre we parse all the cells
      level_bottom = level_top;    //The top of the previous level we considered is the bottom of the current level
      level_top    = nstack;       //The new top is the end of the stack (last cell added from the previous level)
      for(int si=level_bottom;si<level_top;si++){
        const auto c = stack[si];
        //Load donating neighbours of focal cell into the stack
        for(int k=0;k<ndon[c];k++){     
          const auto n = donor[8*c+k];
          stack[nstack++] = n;
        }
      }

      levels[nlevel++] = nstack; //Start a new level
    }

    //End condition for the loop places two identical entries
    //at the end of the stack. Remove one.
    nlevel--;

    assert(levels[nlevel-1]==nstack);
  }













#endif
