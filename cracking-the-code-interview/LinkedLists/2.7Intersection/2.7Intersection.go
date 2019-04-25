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
