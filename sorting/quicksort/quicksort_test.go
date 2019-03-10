package main

import "testing"

// TestQuickSort tests the manual implementation of quicksort
func TestQuickSort(t *testing.T) {
	cases := [][]int{
		[]int{5, 8, 7, 6, 4, 3, 2, 1},
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		[]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	for _, arr := range cases {
		QuickSort(&arr, 0, len(arr)-1)

		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				t.Errorf("Array is not sorted: %d is bigger than %d  ----- sortedArray: %+v", arr[i-1], arr[i], arr)
			}
		}
	}
}
