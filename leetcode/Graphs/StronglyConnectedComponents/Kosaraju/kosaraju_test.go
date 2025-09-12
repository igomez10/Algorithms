package main

import (
	"sort"
	"strconv"
	"testing"
)

func TestKosaraju(t *testing.T) {
	type TestCase struct {
		numNodes int
		graph    [][]int
		expected [][]int
	}

	cases := []TestCase{
		{
			numNodes: 4,
			graph: [][]int{
				{0, 1},
				{1, 2},
				{2, 0},
				{2, 3},
			},
			expected: [][]int{
				{0, 1, 2},
				{3},
			},
		},
	}

	for idx, c := range cases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			ans := Kosaraju(c.numNodes, c.graph)
			if len(ans) != len(c.expected) {
				t.Fatalf("expected %d connected components, got %d", len(c.expected), len(ans))
			}

			for i := range ans {
				sort.Ints(ans[i])
			}

			for i := range ans {
				for j := range ans[i] {
					if ans[i][j] != c.expected[i][j] {
						t.Fatal(ans, c.expected)
					}
				}
			}
		})
	}
}
