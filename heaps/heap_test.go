package main

import "testing"

func TestIsSorted(t *testing.T) {
	goodValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if ok := isSorted(goodValues); !ok {
		m := "isSorted is failing a good value for a binary heap: %+v"
		t.Error(m, goodValues)
	}

	badValues := []int{1, 2, 9, 3, 4, 5, 6, 7, 8}

	if isSorted(badValues) {
		m := "isSorted is passing a bad value for a binary heap: %+v"
		t.Error(m, badValues)
	}
}

func TestInsert(t *testing.T) {
	aHeap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	initialLength := len(aHeap)
	insert(&aHeap, 10)
	if len(aHeap) < initialLength {
		m := "New value %d should be bigger than old value %d"
		t.Error(m, len(aHeap), initialLength)
	}
}
