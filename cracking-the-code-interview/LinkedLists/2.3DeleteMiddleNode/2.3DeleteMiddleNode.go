package main

type SinglyLinkedListNode struct {
	Value int
	Next  *SinglyLinkedListNode
}

func main() {

}

func deleteNode(head **SinglyLinkedListNode, index int) *SinglyLinkedListNode {

	var old *SinglyLinkedListNode
	if index == 0 {
		old, *head = *head, (*head).Next
		return old
	}

	var counter int
	iterativeHead := *head
	for iterativeHead != nil && counter < index-1 {
		iterativeHead = iterativeHead.Next
		counter++
	}

	old = iterativeHead.Next
	(*iterativeHead).Next = iterativeHead.Next.Next
	return old
}

func linkedListFromArray(arr []int) *SinglyLinkedListNode {
	head := SinglyLinkedListNode{Value: arr[0]}
	iterativeHead := &head
	for i := 1; i < len(arr); i++ {
		newNode := SinglyLinkedListNode{Value: arr[i]}
		iterativeHead.Next = &newNode
		iterativeHead = iterativeHead.Next
	}
	return &head
}
