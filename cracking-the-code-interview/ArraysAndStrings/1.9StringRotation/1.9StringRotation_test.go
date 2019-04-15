package main

import (
	"testing"
)

func TestIsRotation(t *testing.T) {

	testCases := [][]string{
		[]string{"Hello", "loHel"},
		[]string{"Hello", "llos"},
	}

	expectedResult := []bool{
		true,
		false,
	}

	for i := range testCases {
		if isRotation(testCases[i][0], testCases[i][1]) != expectedResult[i] {
			t.Errorf("Expected %t for %s %s rotation", expectedResult[i], testCases[i][1], testCases[i][0])
		}
	}

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
	}

	for i := range notSubstrings {
		t.Log("Processing", notSubstrings[i])
		isSub := isSubstring(SuperString, notSubstrings[i])
		if isSub {
			t.Errorf("Expected %s not a substring of %s but got %t", notSubstrings[i], SuperString, isSub)
		}
	}
}
