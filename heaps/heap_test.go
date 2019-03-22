package main

import "testing"

func TestDeleteMin(t *testing.T) {

	values := []int{-1, 1, 2, 3, 4, 5, 6, 7}
	deleteMin(&values)
	expectedArray := []int{-1, 2, 4, 3, 7, 5, 6}
	for i := 0; i < len(values); i++ {
		if values[i] != expectedArray[i] {
			t.Errorf("Expected %d got %d\n", expectedArray[i], values[i])
		}
	}

}

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

func sameSizeArray(t *testing.T, arrA, arrB []int) bool {
	sameSize := true
	if len(arrA) != len(arrB) {
		sameSize = false
		m := `Length of array is not the same: %d was %d`
		t.Errorf(m, len(arrA), len(arrB))
	}
	return sameSize
}

func sameElementsInArray(t *testing.T, arrA, arrB []int) bool {
	hasSameElements := true
	if sameSizeArray(t, arrA, arrB) {
		for i := 0; i < len(arrA); i++ {
			if arrA[i] != arrB[i] {
				hasSameElements = false
				t.Errorf("Values changed in balanced heap: %d, %d", arrA[i], arrB[i])
				break
			}
		}
	}
	return hasSameElements
}

func TestSink(t *testing.T) {
	goodHeap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	copyOfHeap := goodHeap
	//try swim with each possible value in heap
	for i := 0; i < len(goodHeap); i++ {
		sink(&goodHeap, i)
	}

	sameElementsInArray(t, goodHeap, copyOfHeap)

	badValues := []int{1, 3, 4, 5, 6, 7, 8, 3}
	swim(&badValues, len(badValues)-1)
	if !isSorted(badValues) {
		t.Errorf("Heap %+v should be sorted after swim operation", badValues)
	}
}

func TestSwim(t *testing.T) {
	goodHeap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	copyOfHeap := goodHeap
	//try swim with each possible value in heap
	for i := 0; i < len(goodHeap); i++ {
		swim(&goodHeap, i)
	}

	sameElementsInArray(t, goodHeap, copyOfHeap)

	badValues := []int{1, 3, 4, 5, 6, 7, 8, 3}
	swim(&badValues, len(badValues)-1)
	if !isSorted(badValues) {
		t.Errorf("Heap %+v should be sorted after swim operation", badValues)
	}
}
