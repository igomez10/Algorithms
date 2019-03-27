package main

import (
	"testing"
)

func TestIsPermutationOfPalindrome(t *testing.T) {
	s1 := "abcdcbao"
	if isPermutationOfPalindrome(s1) {
		t.Error("Expected true got false")
	}
	s2 := "abcdcba"
	if !isPermutationOfPalindrome(s2) {
		t.Error("Expected true got false")
	}

}
