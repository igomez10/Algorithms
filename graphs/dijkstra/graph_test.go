package dijkstra

import (
	"testing"
)

func TestBuildGraph(t *testing.T) {

	input := [][3]int{
		{1, 2, 7},
		{1, 3, 9},
		{1, 6, 14},
		{2, 3, 10},
		{2, 4, 15},
		{3, 4, 11},
		{3, 6, 2},
		{4, 5, 6},
		{5, 6, 9},
	}

	graph := BuildGraphFromEdges(input)

	if len(graph) != 6 {
		t.Errorf("unexpected graph length, got %d", len(graph))
	}

	node1 := graph[1]
	if len(node1.Neighbors) != 3 {
		t.Errorf("unexpected node1 neighbors length, got %d", len(node1.Neighbors))
	}

	node2 := graph[2]
	if len(node2.Neighbors) != 3 {
		t.Errorf("unexpected node2 neighbors length, got %d", len(node2.Neighbors))
	}

	node3 := graph[3]
	if len(node3.Neighbors) != 4 {
		t.Errorf("unexpected node3 neighbors length, got %d", len(node3.Neighbors))
	}

	node4 := graph[4]
	if len(node4.Neighbors) != 3 {
		t.Errorf("unexpected node4 neighbors length, got %d", len(node4.Neighbors))
	}

	node5 := graph[5]
	if len(node5.Neighbors) != 2 {
		t.Errorf("unexpected node5 neighbors length, got %d", len(node5.Neighbors))
	}

	node6 := graph[6]
	if len(node6.Neighbors) != 3 {
		t.Errorf("unexpected node6 neighbors length, got %d", len(node6.Neighbors))
	}
}
