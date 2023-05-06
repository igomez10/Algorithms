package main

import (
	"testing"
)

func TestShortestWordEditPath(t *testing.T) {

	testcases := []struct {
		source   string
		target   string
		words    []string
		expected int
	}{
		{
			source:   "bit",
			target:   "dog",
			words:    []string{"but", "put", "big", "pot", "pog", "dog", "lot"},
			expected: 5,
		},
		{
			source:   "no",
			target:   "go",
			words:    []string{"to"},
			expected: -1,
		},
		{
			source:   "bit",
			target:   "pog",
			words:    []string{"but", "put", "big", "pot", "pog", "pig", "dog", "lot"},
			expected: 3,
		},
		{
			source:   "aa",
			target:   "bb",
			words:    []string{"ab", "bb"},
			expected: 2,
		},
		{
			source:   "abc",
			target:   "ab",
			words:    []string{"abc", "ab"},
			expected: -1,
		},
		{
			source:   "aa",
			target:   "bbb",
			words:    []string{"ab", "bb"},
			expected: -1,
		},
	}

	for _, tc := range testcases {
		t.Run("source: "+tc.source+" target: "+tc.target, func(t *testing.T) {
			if shortestPath(tc.source, tc.target, tc.words) != tc.expected {
				t.Error("Expected", tc.expected, "got", shortestPath(tc.source, tc.target, tc.words))
			}
		})
	}

	// source := "bit"
	// target := "dog"
	// words := []string{"but", "put", "big", "pot", "pog", "dog", "lot"}
	// expected := 5
	// if shortestPath(source, target, words) != expected {
	// 	t.Error("Expected", expected, "got", shortestPath(source, target, words))
	// }

	// source = "no"
	// target = "go"
	// words = []string{"to"}
	// expected = -1
	// if shortestPath(source, target, words) != expected {
	// 	t.Error("Expected", expected, "got", shortestPath(source, target, words))
	// }

	// source = "bit"
	// target = "pog"
	// words = []string{"but", "put", "big", "pot", "pog", "pig", "dog", "lot"}
	// expected = 3
	// if shortestPath(source, target, words) != expected {
	// 	t.Error("Expected", expected, "got", shortestPath(source, target, words))
	// }

	// source = "aa"
	// target = "bb"
	// words = []string{"ab", "bb"}
	// expected = 2
	// if shortestPath(source, target, words) != expected {
	// 	t.Error("Expected", expected, "got", shortestPath(source, target, words))
	// }

	// source = "abc"
	// target = "ab"
	// words = []string{"abc", "ab"}
	// expected = -1
	// if shortestPath(source, target, words) != expected {
	// 	t.Error("Expected", expected, "got", shortestPath(source, target, words))
	// }

	// source = "aa"
	// target = "bbb"
	// words = []string{"ab", "bb"}
	// expected = -1
	// if shortestPath(source, target, words) != expected {
	// 	t.Error("Expected", expected, "got", shortestPath(source, target, words))
	// }

}
