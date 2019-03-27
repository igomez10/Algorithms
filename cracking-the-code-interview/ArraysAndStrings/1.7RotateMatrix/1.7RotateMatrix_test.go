package main

import (
	"testing"
)

func TestRotateMatrix(t *testing.T) {

	initialMatrix := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15},
		[]int{16, 17, 18, 19, 20},
		[]int{21, 22, 23, 24, 25},
	}

	expectedMatrix := [][]int{
		[]int{21, 16, 11, 6, 1},
		[]int{22, 17, 12, 7, 2},
		[]int{23, 18, 13, 8, 3},
		[]int{24, 19, 14, 9, 4},
		[]int{25, 20, 15, 10, 5},
	}

	RotateMatrix(&initialMatrix)

	for i := range initialMatrix {
		for j := range initialMatrix {
			if initialMatrix[i][j] != expectedMatrix[i][j] {
				expected := matrixToString(expectedMatrix)
				result := matrixToString(initialMatrix)
				t.Errorf("\nExpected \n%s \nBut got:\n\n%s \n", expected, result)
				return
			}
		}
	}
}
