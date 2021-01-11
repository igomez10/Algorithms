package main

import "testing"

func TestMaxProfit(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	expected := 5

	if maxProfit(input) != expected {
		t.Fatal(maxProfit(input))
	}
}
