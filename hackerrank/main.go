package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	initialArray := make([]int32, n)

	for i := 0; i < len(queries); i++ {
		initialArray[queries[i][0]] += queries[i][2]
		initialArray[queries[i][1]] -= queries[i][2]
	}

	maximum := int32(0)
	currentMaximum := int32(0)

	for i := 0; i < len(initialArray); i++ {
		currentMaximum += initialArray[i]
		if currentMaximum > maximum {
			maximum = currentMaximum
		}
	}
	return int64(maximum)
}
