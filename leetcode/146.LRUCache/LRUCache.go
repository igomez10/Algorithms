package main

import (
	"container/heap"
)

type Node struct {
	Key      int
	Value    int
	Position int // position in the heap
	LastSeen int // last call
}

type LRUCache struct {
	Heap        *Nodes
	Cache       map[int]*Node
	Capacity    int
	CurrentCall int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Heap:        &Nodes{},
		Cache:       map[int]*Node{},
		Capacity:    capacity,
		CurrentCall: 0,
	}

	heap.Init(lru.Heap)

	return lru
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.Cache[key]
	if !exist {
		return -1
	}
	this.CurrentCall++
	node.LastSeen = this.CurrentCall
	heap.Fix(this.Heap, node.Position)

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	this.CurrentCall++
	if node, exist := this.Cache[key]; exist {
		node.Value = value
		node.LastSeen = this.CurrentCall
		heap.Fix(this.Heap, node.Position)
		return
	}

	if this.Heap.Len() == this.Capacity {
		removed := heap.Pop(this.Heap)
		delete(this.Cache, removed.(*Node).Key)
	}

	node := &Node{Key: key, Value: value, LastSeen: this.CurrentCall}
	heap.Push(this.Heap, node)
	this.Cache[node.Key] = node
}

type Nodes []*Node

func (n Nodes) Less(i, j int) bool {
	return n[i].LastSeen < n[j].LastSeen
}

func (n Nodes) Swap(i, j int) {
	n[i].Position, n[j].Position = n[j].Position, n[i].Position
	n[i], n[j] = n[j], n[i]
}

func (n Nodes) Len() int {
	return len(n)
}

func (n *Nodes) Push(element interface{}) {
	node := element.(*Node)
	node.Position = len(*n)
	*n = append(*n, node)
}

func (n *Nodes) Pop() interface{} {
	old := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return old
}
