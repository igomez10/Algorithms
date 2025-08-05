package searchrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	type testcases struct {
		nums   []int
		target int
		result int
	}

	tests := []testcases{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1, 3}, 0, -1},
		{[]int{1, 3}, 3, 1},
		{[]int{1, 3}, 1, 0},
		{[]int{1, 3, 5}, 5, 2},
	}

	for _, test := range tests {
		result := search(test.nums, test.target)
		if result != test.result {
			t.Errorf("For %v, expected %d but got %d", test.nums, test.result, result)
		}
	}
}
