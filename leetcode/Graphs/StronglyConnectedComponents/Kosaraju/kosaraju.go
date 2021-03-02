package main

import "container/list"

func Kosaraju(numNodes int, edges [][]int) [][]int {
	graph := buildGraph(numNodes, edges)
	visited := make([]bool, numNodes)
	L := list.New()
	//2. For each node u of the graph, Visit(u) recursively:
	for u := range visited {
		Visit(u, graph, visited, L)
	}

	components := map[int]int{}

	// 3. For each element u in L, Assign(u, u) where Assign(u, root) is:
	for L.Len() > 0 {
		u := L.Remove(L.Front()).(int)
		Assign(u, u, components)
	}
}

func Assign(u, root int, components map[int]int) {
	// If u has not been assigned to a component:
}

func Visit(u int, graph [][]int, visited []bool, L *list.List) {
	if visited[u] {
		return
	}

	// If u is unvisited:
	// 		1. Mark u as visited

	visited[u] = true

	// 2. For each out-neighbor v of u , Visit(v)
	for v := range graph[u] {
		isOutNeighbor := graph[u][v] == 1
		if isOutNeighbor {
			Visit(v, graph, visited, L)
		}

	}

	// 3. Prepend u to L
	L.PushFront(u)
}

/*
	1. For each node u of the graph, mark u as unvisited
	2. For each node u of the graph, Visit(u) recursively:
		If u is unvisited:
			1. Mark u as visited
			2. For each out-neighbor v of u , Visit(v)
			3. Prepend u to L
	3. For each element u in L, Assign(u, u) where Assign(u, root) is:
		If u has not been assigned to a component:
			1. Assign u as belonging to the component whose root is root
			2. For each in-neighbor v of u , Assign(v, root)
*/

func buildGraph(numNodes int, edges [][]int) [][]int {
	graph := make([][]int, numNodes)
	for i := range graph {
		graph[i] = make([]int, numNodes)
	}

	for i := range edges {
		graph[edges[i][0]][edges[i][1]] = 1
	}

	return graph
}
