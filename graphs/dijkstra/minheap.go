package dijkstra

import "container/heap"

type MinHeapNode struct {
	key   string
	value int
}

// MinHeap is a min-heap of integers.
type MinHeap []MinHeapNode

// Len implements heap.Interface.
func (m *MinHeap) Len() int {
	return len(*m)
}

// Less implements heap.Interface.
func (m *MinHeap) Less(i int, j int) bool {
	return (*m)[i].value < (*m)[j].value
}

// Pop implements heap.Interface.
func (m *MinHeap) Pop() any {
	popped := (*m)[len(*m)-1]
	*m = (*m)[0 : len(*m)-1]
	return popped
}

// Push implements heap.Interface.
func (m *MinHeap) Push(x any) {
	*m = append((*m), x.(MinHeapNode))
}

// Swap implements heap.Interface.
func (m *MinHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

var m heap.Interface = &MinHeap{}
