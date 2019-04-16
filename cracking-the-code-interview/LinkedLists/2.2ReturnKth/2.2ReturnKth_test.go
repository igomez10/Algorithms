package main

import "testing"

func TestLinkedListFromArray(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	head := LinkedListFromArray(values)
	var counter int
	for iterativeHead := head; iterativeHead != nil; iterativeHead, counter = iterativeHead.Next, counter+1 {
		if iterativeHead.Value != values[counter] {
			t.Errorf("Unexpected value in index %d, expected %d - got %d", counter, values[counter], iterativeHead.Value)
		}
	}
}

func TestReturnKth(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	head := LinkedListFromArray(values)
	for i := 0; i < len(values); i++ {
		if returnKth(head, i).Value != values[len(values)-i-1] {
			t.Errorf("Unexpected value in index %d, expected %d got %d", i, values[i], returnKth(head, i).Value)
		}
	}
}
