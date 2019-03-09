package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 8, 7, 6, 4, 3, 2, 1}
	fmt.Printf("Initial Array: %+v\n", arr)

	sort(&arr, 0, len(arr)-1)

	fmt.Printf("Sorted Array: %+v\n", arr)
}

func sort(arr *[]int, lo, hi int) {
	if lo >= hi {
		return
	}

	finalIndexForFirstE := partition(arr, lo, hi)
	//first element of the array is already placed in the final position

	sort(arr, lo, finalIndexForFirstE-1) // the first if in sort() will handle any out of bounds from the -1 +1
	sort(arr, finalIndexForFirstE+1, hi)
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

	pivot := (*arr)[lo]
	smallPointer := lo + 1
	bigPointer := hi

	for smallPointer < bigPointer {
		for (*arr)[smallPointer] <= pivot && smallPointer < hi {
			smallPointer++
		}

		for (*arr)[bigPointer] >= pivot && bigPointer > lo {
			bigPointer--
		}

		(*arr)[bigPointer], (*arr)[smallPointer] = (*arr)[smallPointer], (*arr)[bigPointer]
	}

	(*arr)[smallPointer], (*arr)[lo] = (*arr)[lo], (*arr)[smallPointer]

	return bigPointer
}
