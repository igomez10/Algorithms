package main

import "fmt"

func createTargetArray(nums []int, indices []int) []int {
	targetArray := make([]int, len(nums))
	for i, index := range indices {
		copy(targetArray[index+1:i+1], targetArray[index:i])
		targetArray[index] = nums[i]
	}
	return targetArray
}

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}
