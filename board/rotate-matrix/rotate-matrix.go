package main

import "fmt"

func main() {
	arr := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	// fmt.Println(rotateMatrix(arr))

	printMatrix(rotateLayer(arr, 0))

	// [[7,4,1],
	// [8,5,2],
	// [9,6,3]]

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

func rotateLayer(matrix [][]int, layer int) [][]int {

	for i := 0; i < len(matrix)-1; i++ {
		initial := matrix[layer][i]
		tmp := initial

		matrix[i][len(matrix)-1-layer], tmp = tmp, matrix[i][len(matrix)-1-layer]

		matrix[len(matrix)-1-layer][len(matrix)-1-i], tmp = tmp, matrix[len(matrix)-1-layer][len(matrix)-1-i]

		matrix[len(matrix)-1-i][layer], tmp = tmp, matrix[len(matrix)-1-i][layer]

		matrix[layer][i] = tmp

	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println()
}
