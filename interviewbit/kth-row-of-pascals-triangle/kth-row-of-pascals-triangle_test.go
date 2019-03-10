package main

import "fmt"
import "testing"

func testGenerateNextRow(t *testing.T) {
	initialRow := []int{1, 5, 10, 10, 5, 1}
	expectedRow := []int{1, 6, 15, 20, 15, 6, 1}

	nextRow := GenerateNextRow(initialRow)
	fmt.Println(nextRow)
	for i := 0; i < len(nextRow); i++ {
		if nextRow[i] != expectedRow[i] {
			fmt.Println("FAILED")
			t.Errorf("Error")
		} else {
			fmt.Println("SUCCESS")
		}
	}
}
