package main

import (
	"strconv"
	"testing"
)

func TestEditDistance(t *testing.T) {
	type TestCase struct {
		source   string
		target   string
		expected int
	}

	cases := []TestCase{
		{
			source:   "horse",
			target:   "ros",
			expected: 3,
		},
		{
			source:   "intention",
			target:   "execution",
			expected: 5,
		},
		{
			source:   "a",
			target:   "a",
			expected: 0,
		},
		{
			source:   "a",
			target:   "b",
			expected: 1,
		},
		{
			source:   "plasma",
			target:   "altruism",
			expected: 6,
		},
	}

	for i := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if val := minDistance(cases[i].source, cases[i].target); val != cases[i].expected {
				t.Error(cases[i], val)
			}
		})
	}
}
