package main

import (
	"strconv"
	"testing"
)

func TestShoppingOptions(t *testing.T) {
	type TestCase struct {
		jeans    []int
		shoes    []int
		skirts   []int
		tops     []int
		budget   int
		expected int
	}

	cases := []TestCase{
		{
			jeans:    []int{2, 3},
			shoes:    []int{4},
			skirts:   []int{2, 3},
			tops:     []int{1, 2},
			budget:   10,
			expected: 6,
		},
	}

	for idx, c := range cases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			ans := getShoppingOptions(c.jeans, c.shoes, c.skirts, c.tops, c.budget)
			if ans != c.expected {
				t.Errorf("expected %d got %d", c.expected, ans)
			}
		})
	}
}
