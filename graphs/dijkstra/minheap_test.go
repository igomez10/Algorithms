package dijkstra

import (
	"container/heap"
	"testing"
)

func TestMinHeap(t *testing.T) {
	m := MinHeap{}
	heap.Init(&m)
	heap.Push(&m, MinHeapNode{value: 1, key: "1"})
	heap.Push(&m, MinHeapNode{value: 10, key: "10"})
	heap.Push(&m, MinHeapNode{value: 5, key: "5"})
	heap.Push(&m, MinHeapNode{value: -4, key: "-4"})
	heap.Push(&m, MinHeapNode{value: 5, key: "5"})
	heap.Push(&m, MinHeapNode{value: 5, key: "5"})

	if m[0].value != -4 {
		t.Errorf("unexpected value at 0")
	}

	popped := heap.Pop(&m)
	if popped.(MinHeapNode).value != -4 {
		t.Errorf("unexpected element popped")
	}

	if m[0].value != 1 {
		t.Errorf("unexpected value at 0")
	}

	heap.Push(&m, MinHeapNode{key: "", value: -10})

	if m[0].value != -10 {
		t.Errorf("unexpected value at 0")
	}
}
