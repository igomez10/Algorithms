package main

import "testing"

var demoValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var demoHead = linkedListFromArray(demoValues)

func TestLinkedListFromArray(t *testing.T) {

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head := linkedListFromArray(values)
	var counter int

	for head != nil {

		if values[counter] != head.Value {
			t.Errorf("Unexpected value in linked list, expected %d - got %d", values[counter], head.Value)
		}

		head = head.Next
		counter++
	}

}

func TestDeleteNode(t *testing.T) {
	head := demoHead

	firstDeletion := deleteNode(&head, 0)
	if firstDeletion.Value != 1 {
		t.Errorf("Deleted wrong value, expected 0, got %d", firstDeletion.Value)
	}

	if head.Value != 2 {
		t.Errorf("Expected 1 as first element after deletion - got %d ", head.Value)
	}

	secondDeletion := deleteNode(&head, 1)
	if secondDeletion.Value != 3 {
		t.Errorf("Deleted wrong value, expected 3, got %d", secondDeletion.Value)
	}
	if head.Next.Value != 4 {
		t.Errorf("Expected %d as second element after deletion - got %d", 4, head.Next.Value)
	}

	lastDeletion := deleteNode(&head, 6)
	if lastDeletion.Value != 9 {
		t.Errorf("Deleted wrong value, expected 3, got %d", lastDeletion.Value)
	}
}
