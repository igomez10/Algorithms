package main

import "testing"

func TestZeroMatrix(t *testing.T) {
	initialMatrix := [][]int{
		[]int{1, 2, 0},
		[]int{0, 3, 4},
		[]int{5, 6, 7},
	}

	expectedMatrix := [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 6, 0},
	}

	ZeroMatrix(&initialMatrix)

	for i := range initialMatrix {
		for j := range initialMatrix {
			if initialMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf("Expected %d in %d,%d - got %d", expectedMatrix[i][j], i, j, initialMatrix[i][j])
			}
		}
	}
}
