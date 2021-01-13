package main

import "testing"

func TestWordBreak(t *testing.T) {
	type TestCase struct {
		word       string
		dictionary []string
		expected   bool
	}
	cases := []TestCase{
		{
			word:       "leetcode",
			dictionary: []string{"leet", "code"},
			expected:   true,
		},
		{
			word:       "applepenapple",
			dictionary: []string{"apple", "pen"},
			expected:   true,
		},
		{
			word:       "catsandog",
			dictionary: []string{"cats", "dog", "sand", "and", "cat"},
			expected:   false,
		},
	}

	for c := range cases {
		if out := wordBreak(cases[c].word, cases[c].dictionary); out != cases[c].expected {
			t.Error(out, cases[c].expected)
		}
	}
}
