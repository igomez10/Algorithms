package main

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	num := 11
	expected := 3
	if val := coinChangeBottomUp(coins, num); val != expected {
		t.Error(val, expected)
	}

	if val := coinChangeTopBottom(coins, num); val != expected {
		t.Error(val, expected)
	}
}
