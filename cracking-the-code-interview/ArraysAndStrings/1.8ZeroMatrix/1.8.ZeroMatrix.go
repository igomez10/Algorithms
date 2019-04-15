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

	markedXs := make([]int, 0, numcols)
	markedYs := make([]int, 0, numrows)

	//find all 0s in matrix
	for Yi := range *matrix {
		for Xi := range (*matrix)[0] {
			if (*matrix)[Yi][Xi] == 0 {
				markedXs = append(markedXs, Xi)
				markedYs = append(markedYs, Yi)
			}
		}
	}

	for affectedRow := range markedYs {
		NullifyRow(matrix, markedYs[affectedRow])
	}

	for affectedCol := range markedXs {
		NullifyCol(matrix, markedXs[affectedCol])
	}

}

func NullifyRow(matrix *[][]int, index int) {
	(*matrix)[index] = make([]int, len((*matrix)[0]))
}

func NullifyCol(matrix *[][]int, index int) {
	for i := range *matrix {
		(*matrix)[i][index] = 0
	}
}
