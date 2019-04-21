package main

import (
	"fmt"
	"strconv"
	"strings"
)

// LinkedListNode is a node in a singly linked list, only stores reference to next value
type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func main() {
	values := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	head := LLFromArray(values)

	partition(&head, 5)

	fmt.Println(head.ToString())
}

// LLFromArray receives an array of values and returns a reference to a linked list node
// as the first element from the array
func LLFromArray(values []int) *LinkedListNode {
	firstElement := LinkedListNode{value: values[0]}
	head := &firstElement
	for i := 1; i < len(values); i++ {
		head.next = &LinkedListNode{value: values[i]}
		head = head.next
	}
	return &firstElement
}

// ToString returns a string formatted version of the linked list
func (l *LinkedListNode) ToString() string {
	builder := strings.Builder{}

	for l != nil {
		builder.WriteString(strconv.Itoa(l.value))
		l = l.next
	}
	return builder.String()
}

func partition(head **LinkedListNode, value int) {
	var smallerHead *LinkedListNode
	var smallerIterative *LinkedListNode

	var biggerHead *LinkedListNode
	var biggerIterative *LinkedListNode

	var equalsHead *LinkedListNode
	var equalsIterative *LinkedListNode

	iterativeHead := *head

	for iterativeHead != nil {

		if (*iterativeHead).value < value {

			if smallerHead == nil {
				smallerHead = iterativeHead
				smallerIterative = iterativeHead

			} else {
				smallerIterative.next = iterativeHead
				smallerIterative = smallerIterative.next
			}

		} else if (*iterativeHead).value > value {

			if biggerHead == nil {
				biggerHead = iterativeHead
				biggerIterative = iterativeHead

			} else {
				biggerIterative.next = iterativeHead
				biggerIterative = biggerIterative.next
			}

		} else {

			if equalsHead == nil {
				equalsHead = iterativeHead
				equalsIterative = iterativeHead

			} else {
				equalsIterative.next = iterativeHead
				equalsIterative = equalsIterative.next
			}
		}
		iterativeHead = iterativeHead.next
	}

	smallerIterative.next = equalsHead
	equalsIterative.next = biggerHead
	biggerIterative.next = nil

	head = &smallerHead
}
