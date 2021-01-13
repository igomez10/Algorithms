package main

import (
	"strconv"
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	type TestCase struct {
		arr      [][]byte
		expected int
	}

	cases := []TestCase{
		{
			arr: [][]byte{
				{'0', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '1'},
				{'0', '1', '1', '1', '1'},
				{'0', '1', '1', '1', '1'},
				{'0', '0', '1', '1', '1'},
			},
			expected: 9,
		},
		{
			arr: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 4,
		},
	}

	for i := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if res := maximalSquare(cases[i].arr); res != cases[i].expected {
				t.Error(res, cases[i].expected)
			}
		})
	}
}
