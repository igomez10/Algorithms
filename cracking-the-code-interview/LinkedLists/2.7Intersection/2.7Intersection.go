package main

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func main() {}

func Intersection(headA, headB *LinkedListNode) *LinkedListNode {

	if headA == nil || headB == nil {
		return nil
	}

	iterativeHeadA := headA
	iterativeHeadB := headB

	lengthA := 1
	lengthB := 1

	for iterativeHeadA.next != nil {
		lengthA++
		iterativeHeadA = iterativeHeadA.next
	}

	for iterativeHeadB.next != nil {
		lengthB++
		iterativeHeadB = iterativeHeadB.next
	}

	if iterativeHeadB != iterativeHeadA {
		return nil
	}

	longHead := headA
	shortHead := headB

	if lengthA < lengthB {
		longHead, shortHead = shortHead, longHead
	}

	itemsToRemove := lengthA - lengthB
	if itemsToRemove < 0 {
		itemsToRemove = itemsToRemove * -1
	}

	for itemsToRemove != 0 {
		longHead = longHead.next
		itemsToRemove--
	}

	for longHead != nil {
		if longHead == shortHead {
			return longHead
		}
		longHead = longHead.next
		shortHead = shortHead.next
	}
	return nil
}
