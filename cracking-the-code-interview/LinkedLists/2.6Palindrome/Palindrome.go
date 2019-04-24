package main

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func main() {

}

func IsPalindrome(head *LinkedListNode) bool {
	return false
}

func LLFromArray(arr []int) *LinkedListNode {
	initialElement := LinkedListNode{value: arr[0]}
	iterativeHead := &initialElement
	for i := 1; i < len(arr); i++ {
		iterativeHead.next = &LinkedListNode{value: arr[i]}
		iterativeHead = iterativeHead.next
	}

	return &initialElement
}
