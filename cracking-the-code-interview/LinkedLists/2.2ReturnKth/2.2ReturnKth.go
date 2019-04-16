package main

// SinglyLinkedListNode is a node in the linked list
type SinglyLinkedListNode struct {
	Value int
	Next  *SinglyLinkedListNode
}

func main() {

}

func returnKth(head *SinglyLinkedListNode, index int) *SinglyLinkedListNode {
	// get length
	var length int
	for iterativeHead := head; iterativeHead != nil; iterativeHead = iterativeHead.Next {
		length++
	}

	finalHead := head
	for i := 0; i < length-index-1; i++ {
		finalHead = finalHead.Next

	}
	return finalHead
}

// LinkedListFromArray returns a linked list from a given array
func LinkedListFromArray(arr []int) *SinglyLinkedListNode {
	head := SinglyLinkedListNode{Value: arr[0]}
	iterativeHead := &head
	for i := 1; i < len(arr); i++ {
		iterativeHead.Next = &SinglyLinkedListNode{Value: arr[i]}
		iterativeHead = iterativeHead.Next
	}

	return &head
}
