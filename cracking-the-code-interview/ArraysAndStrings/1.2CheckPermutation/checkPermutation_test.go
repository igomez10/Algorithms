package main

import "testing"

func TestCheckPermutation(t *testing.T) {
	isPermutation := []string{
		"abcdefghi",
		"ihgfedcba",
	}
	if !checkPermutation(isPermutation[0], isPermutation[1]) {
		t.Errorf("Expected true got false")
	}

	isNotPermutation := []string{
		"asdfasdflb",
		"falkklasdd",
	}
	if checkPermutation(isNotPermutation[0], isNotPermutation[1]) {
		t.Errorf("Expected false got true")
	}
}
