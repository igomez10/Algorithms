package main

import "testing"

func TestMedianSortedArrays(t *testing.T) {
	type TestCase struct {
		Input  [][]int
		Output float64
	}

	cases := []TestCase{
		{
			Input:  [][]int{{1, 3}, {2}},
			Output: float64(2),
		},
		{
			Input:  [][]int{{1, 2}, {3, 4}},
			Output: float64(2.5),
		},
		{
			Input:  [][]int{{0, 0}, {0, 0}},
			Output: float64(0),
		},
		{
			Input:  [][]int{{}, {1}},
			Output: float64(1),
		},
		{
			Input:  [][]int{{2}, {}},
			Output: float64(2),
		},
		{
			Input:  [][]int{{1, 2}, {1, 2}},
			Output: float64(1.5),
		},
		{
			Input:  [][]int{{3, 4}, {1, 2}},
			Output: float64(2.5),
		},
	}

	for i := range cases {
		t.Run("", func(t *testing.T) {
			if res := findMedianSortedArrays(cases[i].Input[0], cases[i].Input[1]); res != cases[i].Output {
				t.Fatal(cases[i].Input, cases[i].Output, res)
			}
		})
	}

}
