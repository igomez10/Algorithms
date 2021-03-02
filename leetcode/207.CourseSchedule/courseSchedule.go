package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	return !DetectCycle(graph)
}

func DetectCycle(graph [][]bool) bool {
	visited := make([]bool, len(graph))
	stack := make([]bool, len(graph))

	for currOrigin := 0; currOrigin < len(graph); currOrigin++ {
		if detectCycleUtil(currOrigin, graph, &visited, &stack) {
			return true
		}
	}
	return false
}

func detectCycleUtil(origin int, graph [][]bool, visited, stack *[]bool) bool {
	(*visited)[origin] = true
	(*stack)[origin] = true
	defer func() { (*stack)[origin] = false }()
	for target, hasConnection := range graph[origin] {
		if !hasConnection {
			continue
		}
		if (*stack)[target] || (!(*visited)[target] && detectCycleUtil(target, graph, visited, stack)) {
			return true
		}

	}

	return false
}

func buildGraph(vertices int, edges [][]int) [][]bool {
	matrix := make([][]bool, vertices)
	for i := range matrix {
		matrix[i] = make([]bool, vertices)
	}

	for _, e := range edges {
		origin := e[0]
		target := e[1]

		matrix[origin][target] = true

	}

	return matrix
}
