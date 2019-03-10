package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 8, 7, 6, 4, 3, 2, 1}
	fmt.Printf("Initial Array: %+v\n", arr)

	QuickSort(&arr, 0, len(arr)-1)

	fmt.Printf("Sorted Array: %+v\n", arr)
}

// QuickSort sorts the given array using quicksort
func QuickSort(arr *[]int, lo, hi int) {
	if lo >= hi {
		return
	}
	//partition(arr, lo, hi)
	finalIndexForFirstE := partition(arr, lo, hi)
	//first element of the array is already placed in the final position

	QuickSort(arr, lo, finalIndexForFirstE-1) // the first if in sort() will handle any out of bounds from the -1 +1
	QuickSort(arr, finalIndexForFirstE+1, hi)
}

/*
partition return an array where arr[0] is in its final position
decrease hi until hits a smaller num (k < arr[0])
increase lo until hits a bigger num (k > arr[0])
switch arr[lo] with arr[hi]
continue decreasing hi and increasing lo until lo >= hi
when this happens, switch arr[hi] with arr[0], the latter will be in its final position
*/
func partition(arr *[]int, lo, hi int) int {

	smallPointer := lo + 1
	bigPointer := hi

	for {
		for (*arr)[smallPointer] < (*arr)[lo] && smallPointer < hi {
			smallPointer++
		}

		for (*arr)[bigPointer] > (*arr)[lo] && bigPointer > lo {
			bigPointer--
		}

		if smallPointer >= bigPointer {
			break
		}
		(*arr)[bigPointer], (*arr)[smallPointer] = (*arr)[smallPointer], (*arr)[bigPointer]

	}
	(*arr)[lo], (*arr)[bigPointer] = (*arr)[bigPointer], (*arr)[lo]

	return bigPointer
}
