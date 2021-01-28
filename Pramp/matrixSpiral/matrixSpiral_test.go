package main

import (
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}
	expected := []int{1, 2, 3, 4, 5, 10, 15, 20, 19, 18, 17, 16, 11, 6, 7, 8, 9, 14, 13, 12}
	ans := SpiralCopy(input)
	if len(expected) != len(ans) {
		t.Fatal("\n", expected, "\n", ans)
	}

	for i := range expected {
		if expected[i] != ans[i] {
			t.Fatal(expected, "\n", ans)
		}
	}

}
