package main

import (
	"strconv"
	"testing"
)

func TestMaximumTime(t *testing.T) {
	cases := [][]string{
		{
			"?4:5?",
			"14:59",
		},
		{
			"23:5?",
			"23:59",
		},
		{
			"2?:22",
			"23:22",
		},
		{
			"0?:??",
			"09:59",
		},
		{
			"??:??",
			"23:59",
		},
	}

	for c := range cases {
		t.Run(strconv.Itoa(c), func(t *testing.T) {
			maxTime := MaximumTime(cases[c][0])
			if maxTime != cases[c][1] {
				t.Error(cases[c][0], "expected", cases[c][1], "got", maxTime)
			}
		})
	}
}
