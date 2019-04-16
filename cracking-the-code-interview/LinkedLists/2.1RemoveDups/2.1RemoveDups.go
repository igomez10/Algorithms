package main

import (
	"strconv"
	"strings"
)

// SinglyLinkedListNode is a single node in a sinlgy linked list
type SinglyLinkedListNode struct {
	Value int
	Next  *SinglyLinkedListNode
}

func main() {

}

// RemoveDups removes the duplicates elements from a given linked list
func RemoveDups(head *SinglyLinkedListNode) {
	known := map[int]bool{}

	for iterativeHead := head; iterativeHead != nil; iterativeHead = iterativeHead.Next {
		nextElement := iterativeHead.Next
		for nextElement != nil && known[nextElement.Value] {
			nextElement = nextElement.Next
		}

		known[iterativeHead.Value] = true
		iterativeHead.Next = nextElement

	}

}

// NextUnique returns the next unique node, a node whose value is not in the known map
func NextUnique(head *SinglyLinkedListNode, known map[int]bool) *SinglyLinkedListNode {

	iterative := head
	for iterative != nil {
		if known[iterative.Value] {
			iterative = iterative.Next
		} else {
			return iterative
		}
	}
	return nil
}

// SinglyLinkedListFromArray returns a SinglyLinkedList from a given array
func SinglyLinkedListFromArray(values []int) *SinglyLinkedListNode {
	head := &SinglyLinkedListNode{}

	head.Value = values[0]

	iterativeHead := head
	for i := 1; i < len(values); i++ {
		newNode := &SinglyLinkedListNode{
			Value: values[i],
			Next:  nil,
		}
		iterativeHead.Next = newNode
		iterativeHead = iterativeHead.Next
	}

	return head
}

func (s SinglyLinkedListNode) len() int {

	currHead := &s
	var counter int
	for currHead != nil {
		currHead = currHead.Next
		counter++
	}
	return counter
}

// ToString returns the string version of a linked list, good for printing, comparing
func (s SinglyLinkedListNode) ToString() string {
	builder := strings.Builder{}
	currHead := &s
	for currHead != nil {
		builder.WriteString(strconv.Itoa(currHead.Value))
		currHead = currHead.Next
	}
	return builder.String()
}
