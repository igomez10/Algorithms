package main

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func main() {}

func isLoop(head *LinkedListNode) bool {
	fastPointer := head.next
	slowPointer := head

	for fastPointer != nil && fastPointer.next != nil && slowPointer != nil {

		if fastPointer == slowPointer {
			return true
		}

		fastPointer = fastPointer.next.next
		slowPointer = slowPointer.next
	}
	return false
}
