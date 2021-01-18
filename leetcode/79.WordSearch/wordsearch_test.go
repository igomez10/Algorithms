package main

import (
	"strconv"
	"testing"
)

func TestExist(t *testing.T) {
	// board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = 'ABCCED'
	type TestCase struct {
		board    [][]byte
		word     string
		expected bool
	}

	cases := []TestCase{
		// {
		// 	board: [][]byte{
		// 		{'A', 'B', 'C', 'E'},
		// 		{'S', 'F', 'C', 'S'},
		// 		{'A', 'D', 'E', 'E'}},
		// 	word:     "ABCCED",
		// 	expected: true,
		// },
		// {
		// 	board: [][]byte{
		// 		{'C', 'A', 'A'},
		// 		{'A', 'A', 'A'},
		// 		{'B', 'C', 'D'},
		// 	},
		// 	word:     "AAB",
		// 	expected: true,
		// },
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'E', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCESEEEFS",
			expected: true,
		},
	}

	for c := range cases {
		t.Run(strconv.Itoa(c), func(t *testing.T) {
			if res := exist(cases[c].board, cases[c].word); res != cases[c].expected {
				t.Error(cases[c].word, cases[c].expected, res)
			}
		})
	}
}
