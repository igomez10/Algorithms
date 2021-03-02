package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	matrix := buildMatrix(numCourses, prerequisites)
	return !DetectCycles(matrix)
}

func DetectCycles(matrix [][]bool) bool {
	visited := make([]bool, len(matrix))
	stack := make([]bool, len(matrix))

	for origin := range matrix {
		stack[origin] = true
		if hasCycles(origin, matrix, &visited, &stack) {
			return true
		}
		stack[origin] = false
	}
	return false
}

func hasCycles(origin int, matrix [][]bool, visited, stack *[]bool) bool {
	(*visited)[origin] = true
	for target, hasConnection := range matrix[origin] {
		if !hasConnection {
			continue
		}

		if (*stack)[target] {
			return true
		}

		(*stack)[target] = true
		if !(*visited)[target] && hasCycles(target, matrix, visited, stack) {
			return true
		}
		(*stack)[target] = false
	}
	return false
}

func buildMatrix(vertices int, edges [][]int) [][]bool {
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
