package main

import (
	"fmt"
	"testing"
)

func TestDeletionDistance(t *testing.T) {
	cases := [][]string{
		{"dog", "frog", "123"},
		{"some", "some", ""},
		{"some", "thing", "123456789"},
		{"", "", ""},
	}

	for i := range cases {
		if val := DeletionDistance(cases[i][0], cases[i][1]); val != len(cases[i][2]) {
			t.Error("Error", cases[i])
		} else {
			fmt.Println("Success")
		}
	}
}
