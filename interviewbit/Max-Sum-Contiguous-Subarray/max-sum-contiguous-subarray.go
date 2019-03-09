package main

import (
	"fmt"
)

/*
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example:

Given the array [-2,1,-3,4,-1,2,1,-5,4],

the contiguous subarray [4,-1,2,1] has the largest sum = 6.

For this problem, return the maximum sum.


*/

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//arr := []int{-163, -20}

	fmt.Println(maxSubArray(arr))
}
func maxSubArray(arr []int) int {

	maximum := arr[0]

	for i := 0; i < len(arr); i++ {
		maxEndingHere := arr[i]
		if maxEndingHere > maximum {
			maximum = maxEndingHere
		}
		for j := i + 1; j < len(arr) && arr[i] > 0; j++ {
			if maxEndingHere+arr[j] > arr[j] {
				maxEndingHere += arr[j]
			} else {
				maxEndingHere = arr[j]
			}
			maxEndingHere += arr[j]
		}
		if maximum < maxEndingHere {
			maximum = maxEndingHere
		}
	}
	return maximum
}
