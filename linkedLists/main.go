package main

import "fmt"

type singlyLinkedListNode struct {
	data int
	next *singlyLinkedListNode
}

func main() {
	head := &singlyLinkedListNode{data: 1}
	head.next = &singlyLinkedListNode{data: 2}
	head.next.next = &singlyLinkedListNode{data: 3}
	head.next.next.next = &singlyLinkedListNode{data: 4}
	insertToEnd(head, 5)
	printList(head)
}

func insertToEnd(head *singlyLinkedListNode, element int) {
	currHead := head
	for currHead.next != nil {
		currHead = currHead.next
	}
	currHead.next = &singlyLinkedListNode{data: element}
}

func printList(head *singlyLinkedListNode) {
	currHead := head
	for currHead != nil {
		fmt.Println(currHead.data)
		currHead = currHead.next
	}
}
