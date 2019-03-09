package main

import "fmt"

func main() {
	tests := [][]int{
		[]int{1, 2, 0},
		[]int{3, 4, -1, 1},
		[]int{-8, -7, -7},
		[]int{1},
	}
	for i := range tests {
		answer := firstMissingPositive(tests[i])
		fmt.Println(answer)
	}
}

func firstMissingPositive(arr []int) int {

	newArr := make([]bool, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		if arr[i] <= len(arr) && arr[i] >= 0 {
			newArr[arr[i]] = true
		}
	}
	fmt.Println(newArr)
	for i := 1; i < len(newArr); i++ {
		if !newArr[i] {
			return i
		}
	}
	return len(newArr)
}
