package main

import (
	"strconv"
	"testing"
)

func TestBoldTags(t *testing.T) {

	type Testcase struct {
		s        string
		dict     []string
		expected string
	}

	cases := []Testcase{
		{
			s:        "abcxyz123",
			dict:     []string{"abc", "123"},
			expected: "<b>abc</b>xyz<b>123</b>",
		},
		{
			s:        "aaabbcc",
			dict:     []string{"aaa", "aab", "bc"},
			expected: "<b>aaabbc</b>c",
		},
		{
			s:        "aaabbcc",
			dict:     []string{"aaa", "aab", "bc", "aaabbcc"},
			expected: "<b>aaabbcc</b>",
		},
		{
			s:        "aaabbcc",
			dict:     []string{"a", "b", "c"},
			expected: "<b>aaabbcc</b>",
		},
	}

	for i := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res := addBoldTag(cases[i].s, cases[i].dict)
			if res != cases[i].expected {
				t.Error("\nfailed", res, cases[i].expected)
			}
		})
	}

}
