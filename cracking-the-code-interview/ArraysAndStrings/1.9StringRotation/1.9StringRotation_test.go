package main

import (
	"testing"
)

func TestStringRotation(t *testing.T) {
	//	s1 := "Hello"
	//	s2 := "loHel"

}

func TestIsSubstring(t *testing.T) {
	SuperString := "Hello"

	substrings := []string{
		"Hell",
		"llo",
		"o",
		"H",
		"ll",
	}

	for i := range substrings {
		t.Log("Processing", substrings[i])
		isSub := isSubstring(SuperString, substrings[i])
		if !isSub {
			t.Errorf("Expected %s substring of %s but got %t", substrings[i], SuperString, isSub)
		}
	}

	notSubstrings := []string{
		"ABC",
		"QWE",
		"HELLO",
		"Hella",
		"lloi",
		"Hello",
	}

	for i := range notSubstrings {
		t.Log("Processing", notSubstrings[i])
		isSub := isSubstring(SuperString, notSubstrings[i])
		if isSub {
			t.Errorf("Expected %s not a substring of %s but got %t", notSubstrings[i], SuperString, isSub)
		}
	}
}
