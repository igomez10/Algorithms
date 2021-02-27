package main

import "fmt"

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
	Prev *ListNode
}

type LRUCache struct {
	dict     map[int]*ListNode
	head     *ListNode
	tail     *ListNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &ListNode{}
	tail := &ListNode{}

	head.Next = tail
	tail.Prev = head

	cache := LRUCache{
		capacity: capacity,
		dict:     map[int]*ListNode{},
		head:     head,
		tail:     tail,
	}

	return cache
}

func (this *LRUCache) Get(key int) int {
	// fmt.Println(this.dict)
	node, exist := this.dict[key]
	if !exist {
		return -1
	}

	this.MoveToHead(key)

	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	// fmt.Println(this.dict)
	node, exist := this.dict[key]
	if exist {
		node.Val = value // overwrite value
		err := this.MoveToHead(key)
		if err != nil {
			// fmt.Println(err)
		}
		return
	}

	if len(this.dict) >= this.capacity {
		this.DeleteOldest()
	}

	newNode := &ListNode{
		Key:  key,
		Val:  value,
		Next: this.head.Next,
		Prev: this.head,
	}

	this.dict[key] = newNode
	this.head.Next.Prev = newNode
	this.head.Next = newNode
}

func (this *LRUCache) MoveToHead(key int) error {
	// fmt.Println("moving", key, "to head")
	node, exist := this.dict[key]
	if !exist {
		return fmt.Errorf("error moving key to head invalid key %q", key)
	}

	//remove
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next

	//move to head
	node.Prev = this.head
	node.Next = this.head.Next

	this.head.Next.Prev = node
	this.head.Next = node
	return nil
}

func (this *LRUCache) DeleteOldest() {
	oldest := this.tail.Prev
	delete(this.dict, oldest.Key)

	oldest.Prev.Next = oldest.Next
	oldest.Next.Prev = oldest.Prev
}
