package main

import "testing"

func TestSumLL(t *testing.T) {
	nodeA := llnode{value: 7}
	nodeB := llnode{value: 1}
	nodeC := llnode{value: 6}

	nodeD := llnode{value: 5}
	nodeE := llnode{value: 9}
	nodeF := llnode{value: 2}

	nodeA.next = &nodeB
	nodeB.next = &nodeC

	nodeD.next = &nodeE
	nodeE.next = &nodeF

	listA := &nodeA
	listB := &nodeD

	nodeX := llnode{value: 2}
	nodeY := llnode{value: 1}
	nodeZ := llnode{value: 9}

	nodeX.next = &nodeY
	nodeY.next = &nodeZ

	expectedList := &nodeX

	newList := sumLists(listA, listB)

	for newList != nil || expectedList != nil {

		if newList.value != expectedList.value {
			t.Errorf("Expected %d got %d", expectedList.value, newList.value)
		}

		expectedList, newList = expectedList.next, newList.next
	}

}
