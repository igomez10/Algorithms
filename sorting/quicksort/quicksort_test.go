package main

import "testing"

// TestQuickSort tests the manual implementation of quicksort
func TestQuickSort(t *testing.T) {
	arr := []int{5, 8, 7, 6, 4, 3, 2, 1}
	QuickSort(&arr, 0, len(arr)-1)

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			t.Errorf("Array is not sorted: %d is bigger than %d  ----- sortedArray: %+v", arr[i-1], arr[i], arr)
		}
	}
}
