package main

import "fmt"

func gameOfLife(matrix [][]int) {

	newMatrix := make([][]int, len(matrix))
	for i := range newMatrix {
		newMatrix[i] = make([]int, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			live := getLiveNeighbors(matrix, i, j)
			currPosition := matrix[i][j]
			if currPosition == 0 && live == 3 {
				newMatrix[i][j] = 1
			} else if currPosition == 1 && (live == 2 || live == 3) {
				newMatrix[i][j] = 1
			}
		}
	}
	// matrix = [][]int{}
	matrix = newMatrix
	// return matrix
}

func getLiveNeighbors(matrix [][]int, i, j int) int {
	count := 0
	if i > 0 && j > 0 && matrix[i-1][j-1] == 1 {
		count++
	}

	if i > 0 && j < len(matrix[i])-1 && matrix[i-1][j+1] == 1 {
		count++
	}

	if i > 0 && matrix[i-1][j] == 1 {
		count++
	}

	if i < len(matrix)-1 && j > 0 && matrix[i+1][j-1] == 1 {
		count++
	}
	if i < len(matrix)-1 && j < len(matrix[i])-1 && matrix[i+1][j+1] == 1 {
		count++
	}

	if i < len(matrix)-1 && matrix[i+1][j] == 1 {
		count++
	}
	if j > 0 && matrix[i][j-1] == 1 {
		count++
	}
	if j < len(matrix[i])-1 && matrix[i][j+1] == 1 {
		count++
	}

	return count
}

func main() {
	matrix := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	fmt.Println()
	printMatrix(matrix)
	// newMatrix := gameOfLife(matrix)
	fmt.Println()
	// printMatrix(newMatrix)
}

func printMatrix(matrix [][]int) {
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

// [[0,0,0],
// [1,0,1],
// [0,1,1],
// [0,1,0]]
