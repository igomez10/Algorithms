package main

import (
	"fmt"
	"testing"
)

func TestMinJob(t *testing.T) {
	type TestCase struct {
		jobs     []int
		days     int
		expected int
	}
	cases := []TestCase{
		{
			jobs:     []int{6, 5, 4, 3, 2, 1},
			days:     2,
			expected: 7,
		},
		{
			jobs:     []int{9, 9, 9},
			days:     4,
			expected: -1,
		},
		{
			jobs:     []int{1, 1, 1},
			days:     3,
			expected: 3,
		},
		{
			jobs:     []int{7, 1, 7, 1, 7, 1},
			days:     3,
			expected: 15,
		},
		{
			jobs:     []int{11, 111, 22, 222, 33, 333, 44, 444},
			days:     6,
			expected: 843,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%+v  days: %d expected %d", c.jobs, c.days, c.expected), func(t *testing.T) {
			if val := minDifficultyBottomUP(c.jobs, c.days); val != c.expected {
				t.Errorf("expected %d got %d", c.expected, val)
			}
		})
	}
}
