package main

import "fmt"

func main() {
	//	matrix := [][]int{}
	fmt.Println("vim-go")
}

// ZeroMatrix takes the given pointer to matrix of ints and transforms
// any row or column including 0 to a row or col of 0s
func ZeroMatrix(matrix *[][]int) {
	numcols, numrows := len(*matrix), len((*matrix)[0])

	markedXs := make([]bool, numcols)
	markedYs := make([]bool, numrows)

	//find all 0s in matrix
	for Yi := range *matrix {
		for Xi := range (*matrix)[0] {
			if (*matrix)[Yi][Xi] == 0 {
				markedXs[Xi] = true
				markedYs[Yi] = true
			}
		}
	}

	//change all numbers in affected rows, cols
	for Yi := range *matrix {
		for Xi := range (*matrix)[0] {
			if markedXs[Xi] || markedYs[Yi] {
				(*matrix)[Yi][Xi] = 0
			}
		}
	}
}
