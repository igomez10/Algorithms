package main

import (
	"regexp"
	"strings"
)

type SinlgyLinkedListNode struct {
	value int
	next  *SinlgyLinkedListNode
}

type LRUCache struct {
	cache    map[int]int
	capacity int
	MRU      *SinlgyLinkedListNode
}

func Constructor(capacity int) LRUCache {
	newCache := make(map[int]int)
	cap := capacity
	newMRU := SinlgyLinkedListNode{}
	newLRUCache := LRUCache{cache: newCache, capacity: cap, MRU: &newMRU}
	re := regexp.MustCompile(`^\ *[+-]?[0-9]+\.?[0-9]?`)
	found := re.Find
	strings.TrimLeft

	return newLRUCache
}

func (this *LRUCache) Get(key int) int {
	if val, exist := this.cache[key]; exist {
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if this.MRU == nil {
		this.MRU.value = key
	}

	if this.capacity > len(this.cache) {
		// only add if we have space
		this.cache[key] = value

	} else {
		iterativeHead := this.MRU

		for ; iterativeHead.next != nil; iterativeHead = iterativeHead.next {

		}

		delete(this.cache, iterativeHead.value)
		this.cache[key] = value
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {}
