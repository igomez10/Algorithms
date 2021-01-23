package main

import (
	"strconv"
	"testing"
)

func TestDiffStrings(t *testing.T) {
	type TestCase struct {
		source   string
		target   string
		expected []string
	}

	cases := []TestCase{
		{
			source:   "ABCDEFG",
			target:   "ABDFFGH",
			expected: []string{"A", "B", "-C", "D", "-E", "F", "+F", "G", "+H"},
		},
	}

	for i := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := DiffStrings(cases[i].source, cases[i].target)
			for j := range res {
				if res[j] != cases[i].expected[j] {
					t.Error(res)
					t.Error(cases[i].expected)
				}
			}
		})
	}
}
