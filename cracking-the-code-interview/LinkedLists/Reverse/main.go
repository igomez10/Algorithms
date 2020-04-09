package main

import (
	"fmt"
	"strconv"
	"strings"
)

// SinglyLinkedListNode is a single node in a sinlgy linked list
type SinglyLinkedListNode struct {
	Value int
	Next  *SinglyLinkedListNode
}

func main() {
	head := SinglyLinkedListFromArray([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})

	newHead := &SinglyLinkedListNode{}

	for head.Next != nil {
		tmp := head
		head = head.Next

		tmp.Next = newHead.Next
		newHead.Next = tmp
	}

	head.Next = newHead.Next
	newHead = head

	fmt.Println(head.ToString())
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
