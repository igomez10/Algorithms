package main

import "testing"

func TestCountInversions(t *testing.T) {
	input1 := [][]int{
		[]int{1, 3, 2, 0},
		[]int{1},
	}

	countInversions1 := CountInversions(input1[0])
	if countInversions1 != input1[1][0] {
		t.Errorf("Expected %d inversions - got %d", input1[1][0], countInversions1)
	}
}

func TestCountInversionsNaive(t *testing.T) {
	input1 := [][]int{
		[]int{1, 3, 2, 0},
		[]int{4},
	}

	countInversions1 := CountInversionsNaive(input1[0])
	if countInversions1 != input1[1][0] {
		t.Errorf("Expected %d inversions - got %d", input1[1][0], countInversions1)
	}
}
