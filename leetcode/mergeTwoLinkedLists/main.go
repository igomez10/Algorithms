package main

import "fmt"

func main() {
	l1 := LinkedListFromArray([]int{1, 3, 5, 7, 9})
	// l2 := LinkedListFromArray([]int{0, 2, 4, 6, 8})

	// l3 := mergeTwoLists(l1, l2)
	// printList(l3)
	popLeft(l1)
	// fmt.Println(element)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 ListNode, l2 ListNode) *ListNode {

	head := &ListNode{}
	// tail := head

	_, _, smallest := getSmallest(&l1, &l2)
	fmt.Println(smallest)
	// smallest :=
	// for smallest != nil {
	// 	tail.Next = smallest
	// 	tail = tail.Next
	// 	smallest = getSmallest(l1, l2)
	// }

	return head.Next
}

func getSmallest(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	l1.Val += 10
	l1 = l1.Next
	if true {
		return nil, nil, nil
	}

	var tmp *ListNode
	if l1 == nil && l2 == nil {
		return l1, l2, nil

	} else if l1 == nil || l2 == nil {

		if l1 == nil {
			tmp = l1
			l1 = l1.Next
			tmp.Next = nil
			return l1, l2, tmp
		}

		tmp = l2
		l2 = l2.Next
		tmp.Next = nil
		return l1, l2, tmp
	}

	if l1.Val < l2.Val {
		tmp = l1
		l1 = l1.Next
		tmp.Next = nil
		return l1, l2, tmp
	}

	tmp = l2
	l2 = l2.Next
	tmp.Next = nil
	return l1, l2, tmp
}

// SinglyLinkedListFromArray returns a SinglyLinkedList from a given array
func LinkedListFromArray(values []int) *ListNode {
	head := &ListNode{}

	head.Val = values[0]

	iterativeHead := head
	for i := 1; i < len(values); i++ {
		newNode := &ListNode{
			Val:  values[i],
			Next: nil,
		}
		iterativeHead.Next = newNode
		iterativeHead = iterativeHead.Next
	}

	return head
}

func printList(head *ListNode) {
	currHead := head
	for currHead != nil {
		fmt.Println(currHead.Val)
		currHead = currHead.Next
	}
}

func popLeft(head *ListNode) {

	head = head.Next
	head = head.Next
	head = head.Next

}
