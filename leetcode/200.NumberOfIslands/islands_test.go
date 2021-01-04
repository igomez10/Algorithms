package main

import "testing"

func TestIslands(t *testing.T) {

	caseA := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '1'},
	}
	expectedA := 1

	if ans := numIslands(caseA); ans != expectedA {
		t.Errorf("expected %d got %d", expectedA, ans)
	}

}
