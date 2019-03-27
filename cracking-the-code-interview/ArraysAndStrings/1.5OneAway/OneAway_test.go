package main

import (
	"testing"
)

func TestIsOneWay(t *testing.T) {
	A := []string{"pale", "ple"}
	B := []string{"pales", "pale"}
	C := []string{"pale", "bale"}
	D := []string{"pale", "bae"}

	expected := []bool{true, true, true, false}
	input := [][]string{A, B, C, D}
	for i := range input {
		currentInput := input[i]
		if isOneAway(currentInput[0], currentInput[1]) != expected[i] {
			t.Errorf("Expected %t for %+v but got %t", expected[i], currentInput[i], !expected[i])
		}
	}
}
