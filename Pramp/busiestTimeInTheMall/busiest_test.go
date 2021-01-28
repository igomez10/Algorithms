package main

import (
	"strconv"
	"testing"
)

func TestBusiest(t *testing.T) {
	type TestCase struct {
		timestamps [][]int
		expected   int
	}
	cases := []TestCase{
		{
			timestamps: [][]int{
				{1487799425, 14, 1},
				{1487799425, 4, 0},
				{1487799425, 2, 0},
				{1487800378, 10, 1},
				{1487801478, 18, 0},
				{1487801478, 18, 1},
				{1487901013, 1, 0},
				{1487901211, 7, 1},
				{1487901211, 7, 0},
			},
			expected: 1487800378,
		},
		{
			timestamps: [][]int{
				{1487799425, 21, 0},
				{1487799427, 22, 1},
				{1487901318, 7, 0}},
			expected: 1487799427,
		},
		{
			timestamps: [][]int{
				{1487799425, 14, 1},
				{1487799425, 4, 1},
				{1487799425, 2, 1}, // 20
				{1487800378, 10, 1},
				{1487801478, 18, 1}, // 28
				{1487901013, 1, 1},  // 29
				{1487901211, 7, 1},
				{1487901211, 7, 1}, //  43
			},
			expected: 1487901211,
		},
	}

	for c := range cases {
		t.Run(strconv.Itoa(c), func(t *testing.T) {
			if res := FindBusiestPeriod(cases[c].timestamps); res != cases[c].expected {
				t.Error("expected", cases[c].expected, "got", res)
			}
		})
	}

}
