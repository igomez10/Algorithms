package main

import (
	"strconv"
	"testing"
)

func TestSmallestSubstring(t *testing.T) {
	type Testcase struct {
		arr      []string
		str      string
		expected string
	}
	cases := []Testcase{
		{
			arr:      []string{"x", "y", "z"},
			str:      "xyyzyzyx",
			expected: "zyx",
		},
	}

	for i := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := GetShortestUniqueSubstring(cases[i].arr, cases[i].str)
			if res != cases[i].expected {
				t.Error("\n", cases[i].expected, "\n", res)
			}
		})
	}
}
