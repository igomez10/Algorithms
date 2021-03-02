package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type Testcase struct {
		input        []interface{}
		expectedRoot *TreeNode
	}

	cases := []Testcase{
		{
			input:        []interface{}{},
			expectedRoot: nil,
		},
		{
			input: []interface{}{1, 2, 3},
			expectedRoot: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			input: []interface{}{1, 2, 3},
			expectedRoot: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	for i, c := range cases {

		t.Run(strconv.Itoa(i), func(t *testing.T) {

			root := buildTree(c.input)
			valsRoot := BFS(root)
			expectedVals := BFS(c.expectedRoot)
			if len(valsRoot) != len(expectedVals) {
				t.Fatal(valsRoot, expectedVals)
			}

			for i := range valsRoot {
				if valsRoot[i] != expectedVals[i] {
					t.Fatal(valsRoot, expectedVals)
				}
			}
		})
	}
}

func TestBuildGraph(t *testing.T) {
	type Testcase struct {
		vertices  int
		edges     [][]int
		adjacency map[int]map[int]bool
	}

	cases := []Testcase{
		{
			vertices: 3,
			edges:    [][]int{{1, 2}, {2, 3}},
			adjacency: map[int]map[int]bool{
				1: {2: true},
				2: {3: true},
			},
		},
		{
			vertices:  0,
			edges:     [][]int{},
			adjacency: map[int]map[int]bool{},
		},
		{
			vertices: 6,
			edges: [][]int{
				{1, 5},
				{2, 5},
				{3, 5},
				{4, 5},
				{5, 6},
			},
			adjacency: map[int]map[int]bool{
				1: {5: true},
				2: {5: true},
				3: {5: true},
				4: {5: true},
				5: {6: true},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			graph := buildGraph(c.vertices, c.edges)
			for origin, children := range c.adjacency {
				if _, exist := graph[origin]; !exist {
					t.Fatal("expected node does'nt exist")
				}

				if !reflect.DeepEqual(graph[origin].Children, children) {
					t.Fatal(graph[origin].Children, children)
				}
			}
		})
	}
}

func TestPathExist(t *testing.T) {
	type Testcase struct {
		vertices int
		edges    [][]int
		origin   int
		target   int
		expected bool
	}

	cases := []Testcase{
		{
			vertices: 3,
			edges:    [][]int{{1, 2}, {2, 3}},
			origin:   1,
			target:   3,
			expected: true,
		},
		{
			vertices: 6,
			edges: [][]int{
				{1, 5},
				{2, 5},
				{3, 5},
				{4, 5},
				{5, 6},
			},
			origin:   1,
			target:   6,
			expected: true,
		},
		{
			vertices: 6,
			edges: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{5, 6},
			},
			origin:   1,
			target:   6,
			expected: true,
		},
		{
			vertices: 6,
			edges: [][]int{
				{1, 4},
				{2, 4},
				{3, 4},
				{5, 6},
			},
			origin:   1,
			target:   6,
			expected: false,
		},
		{
			vertices: 2,
			edges:    [][]int{},
			origin:   1,
			target:   2,
			expected: false,
		},
		{
			vertices: 3,
			edges: [][]int{
				{1, 1},
				{1, 2},
				{2, 3},
				{3, 1},
			},
			origin:   1,
			target:   2,
			expected: true,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			graph := buildGraph(c.vertices, c.edges)
			exist := pathExist(c.origin, c.target, graph, map[int]bool{})
			if exist != c.expected {
				t.Error("expected", c.expected, "got", exist)
			}
		})
	}
}

func TestDetectCycle(t *testing.T) {
	type Testcase struct {
		vertices int
		edges    [][]int
		hasCycle bool
	}

	cases := []Testcase{
		// {
		// 	vertices: 3,
		// 	edges:    [][]int{{1, 2}, {2, 3}},
		// 	hasCycle: false,
		// },
		// {
		// 	vertices: 6,
		// 	edges: [][]int{
		// 		{1, 5},
		// 		{2, 5},
		// 		{3, 5},
		// 		{4, 5},
		// 		{5, 6},
		// 	},
		// 	hasCycle: false,
		// },
		// {
		// 	vertices: 6,
		// 	edges: [][]int{
		// 		{1, 2},
		// 		{2, 3},
		// 		{3, 4},
		// 		{4, 5},
		// 		{5, 6},
		// 	},
		// 	hasCycle: false,
		// },
		// {
		// 	vertices: 6,
		// 	edges: [][]int{
		// 		{1, 4},
		// 		{2, 4},
		// 		{3, 4},
		// 		{5, 6},
		// 	},
		// 	hasCycle: false,
		// },
		// {
		// 	vertices: 2,
		// 	edges:    [][]int{},
		// 	hasCycle: false,
		// },
		// {
		// 	vertices: 3,
		// 	edges: [][]int{
		// 		{1, 2},
		// 		{2, 3},
		// 		{3, 1},
		// 	},
		// 	hasCycle: true,
		// },
		// {
		// 	vertices: 6,
		// 	edges: [][]int{
		// 		{1, 4},
		// 		{2, 4},
		// 		{3, 4},
		// 		{4, 5},
		// 		{4, 6},
		// 	},
		// 	hasCycle: false,
		// },
		// {
		// 	vertices: 6,
		// 	edges: [][]int{
		// 		{1, 4},
		// 		{2, 4},
		// 		{3, 4},
		// 		{4, 5},
		// 		{4, 6},
		// 		{6, 1},
		// 	},
		// 	hasCycle: true,
		// },
		{
			vertices: 2,
			edges: [][]int{
				{2, 1},
				{1, 2},
			},
			hasCycle: true,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			graph := buildGraph(c.vertices, c.edges)
			exist := DetectCycle(c.vertices, graph)
			if exist != c.hasCycle {
				t.Error("expected", c.hasCycle, "got", exist)
			}
		})
	}
}
