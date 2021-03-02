package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestDFS(t *testing.T) {

	cases := [][]interface{}{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	for i := range cases {
		graph := buildTree(cases[i])
		bfsValue := BFS(graph)
		dfsValue := DFS(graph)
		if len(bfsValue) != len(dfsValue) {
			t.Fatal("unexpected length", len(bfsValue), len(dfsValue))
		}
		sort.Ints(bfsValue)
		sort.Ints(dfsValue)

		if !reflect.DeepEqual(dfsValue, bfsValue) {
			t.Errorf("unexpected values")
		}

	}
}
