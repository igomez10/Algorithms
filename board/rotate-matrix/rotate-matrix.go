package main

import "fmt"

func main() {
	arr := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(rotateMatrix(arr))
}

func rotateMatrix(arr [][]int) [][]int {
	newMatrix := make([][]int, len(arr))
	for i := range newMatrix {
		newMatrix[i] = make([]int, len(arr[0]))
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			newMatrix[j][len(arr)-1-i] = arr[i][j]
		}
	}
	return newMatrix
}
