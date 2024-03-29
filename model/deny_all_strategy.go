package model

import (
	"container/heap"
)

type DenyAllStrategy struct {
	BaseStrategy
	leastReplicatedRank int
	rankCount           RankCount
}

func (strategy *DenyAllStrategy) countBeingRemovedReplicas() int {
	return strategy.countEdgeReplicas()
}

func (strategy *DenyAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()
	rank := rankedStudent.Rank()

	if !jc.IsFull() {
		lrr := strategy.leastReplicatedRank
		if lrr == 0 || rank < lrr {
			heap.Push(pq, rankedStudent)
			jc.DecSpots()
			return true
		}

		if rank == lrr {
			strategy.rankCount[rank] += 1
		}
		return false
	}

	tmp := heap.Pop(pq).(*RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()
	if rank > lastRank {
		return false
	}

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas()
	strategy.rankCount[lastRank], strategy.leastReplicatedRank = count, lastRank
	if count > 0 {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
	}
	for i := 1; i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
		jc.IncSpots()
	}

	return rank < lastRank
}
