package main

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// removes the first item. This method is not supposed to be called directly
func (i *minHeap) Pop() interface{} {
	iDeref := *i
	last := (*i)[len(iDeref)-1]
	*i = iDeref[0 : len(iDeref)-1]
	return last
}

// adds a new item. This method is not supposed to be called directly
func (i *minHeap) Push(x interface{}) {
	*i = append(*i, x.(int))
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// sort intervals based on starting time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// initialize the meeting rooms heap
	ivls := &minHeap{}
	heap.Init(ivls)

	heap.Push(ivls, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < (*ivls)[0] {
			heap.Push(ivls, intervals[i][1])
		} else {
			heap.Pop(ivls)
			heap.Push(ivls, intervals[i][1])
		}
	}

	return ivls.Len()
}
