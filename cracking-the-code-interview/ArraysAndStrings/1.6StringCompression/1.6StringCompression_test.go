package main

import (
	"testing"
)

func TestStringCompression(t *testing.T) {
	first := []string{"aaabbc", "a3b2c1"}
	second := []string{"abbcccddddeeeee", "a1b2c3d4e5"}
	third := []string{"abcdefghi", "abcdefghi"}
	cases := [][]string{first, second, third}

	for i := range cases {

		result := stringCompression(cases[i][0])
		if len(result) != len(cases[i][1]) {
			t.Errorf("Unexpected lenght of compression, expected %s got %s", result, cases[i][1])
		}
		for j := range result {
			if result[j] != cases[i][1][j] {
				t.Errorf("Expected string %s but got %s", result, cases[i][1])
			}
		}
	}

}
