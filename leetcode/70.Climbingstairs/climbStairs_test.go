package main

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := [][]int{
		{2, 2},
		{3, 3},
	}

	for i := range cases {
		if climbStairs(cases[i][0]) != cases[i][1] {
			t.Errorf("expected %d got %d", cases[i][0], cases[i][1])
		}
	}
}
