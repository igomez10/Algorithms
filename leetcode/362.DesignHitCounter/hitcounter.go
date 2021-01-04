package main

import (
	"container/heap"
)

type HitCounter struct {
	TimestampsHeap *Timestamps
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	ts := &Timestamps{}
	heap.Init(ts)

	h := HitCounter{ts}
	return h
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	heap.Push(this.TimestampsHeap, timestamp)
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	for len(*this.TimestampsHeap) > 0 && timestamp-(*this.TimestampsHeap)[0] >= 300 {
		heap.Pop(this.TimestampsHeap)
	}

	return len(*this.TimestampsHeap)
}

type Timestamps []int

func (t Timestamps) Less(i, j int) bool { return t[i] < t[j] }
func (t Timestamps) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Timestamps) Len() int           { return len(t) }
func (t *Timestamps) Pop() interface{} {
	tDeref := *t
	old := tDeref[len(tDeref)-1]
	*t = tDeref[:len(tDeref)-1]
	return old
}
func (t *Timestamps) Push(i interface{}) {
	*t = append(*t, i.(int))
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
