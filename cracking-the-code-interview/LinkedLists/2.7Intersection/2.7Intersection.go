package main

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func main() {}

func Intersection(headA, headB *LinkedListNode) *LinkedListNode {
	iterativeHeadA := headA

	for iterativeHeadA != nil {

		iterativeHeadB := headB
		for iterativeHeadB != nil {
			if iterativeHeadA == iterativeHeadB {
				return iterativeHeadA
			}
			iterativeHeadB = iterativeHeadB.next
		}
		iterativeHeadA = iterativeHeadA.next
	}

	return nil
}

func count(head *LinkedListNode) int {
	if head == nil {
		return 0
	}
	counter := 0
	for head != nil {
		counter++
		head = head.next
	}

	return counter
}
