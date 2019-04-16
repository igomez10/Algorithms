package main

import (
	"testing"
)

func TestCreateSinglyLinkedList(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head := SinglyLinkedListFromArray(values)
	var counter int
	for head == nil {
		if values[counter] != head.Value {
			t.Errorf("Unexpected value in index %d, expected %d got %d ", counter, values[counter], head.Value)
		}
		counter++
		head = head.Next
	}
}

func TestLen(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head := SinglyLinkedListFromArray(values)
	if head.len() != len(values) {
		t.Errorf("Unexpected length of linkedlist, expected %d got %d", len(values), head.len())
		t.Log(head.ToString())
	}
}

func TestRemoveDups(t *testing.T) {
	uniques := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	head := SinglyLinkedListFromArray(uniques)
	RemoveDups(head)
	var counter int

	for iterativeHead := head; iterativeHead != nil; iterativeHead, counter = iterativeHead.Next, counter+1 {
		if uniques[counter] != iterativeHead.Value {
			t.Log(iterativeHead.ToString())
			t.Errorf("Unexpected value in linked list after removing duplicates, expected %d got %d", uniques[counter], iterativeHead.Value)
		}
	}

	duplicated := []int{9, 1, 2, 9, 9, 9, 3, 1, 2, 2, 4}
	duplicatedHead := SinglyLinkedListFromArray(duplicated)

	RemoveDups(duplicatedHead)

	if duplicatedHead.len() != 5 {
		t.Errorf("Not Cleaning enough, length should be %d after removing dups, was %d", 5, duplicatedHead.len())
		t.Log(duplicatedHead.ToString())
	}
}

func TestNextUnique(t *testing.T) {
	values := []int{1, 2, 3, 1, 2, 3, 4}
	head := SinglyLinkedListFromArray(values)

	known := map[int]bool{
		1: true,
		2: true,
		3: true,
	}

	next := NextUnique(head, known)
	if next == nil {
		t.Errorf("Expected next unique 4 but got nil")
		t.FailNow()
	}
	if next.Value != 4 {
		t.Errorf("Expected next unique 4 but got %d", next.Value)
	}

}
