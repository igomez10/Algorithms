package main

import (
	"container/list"
	"fmt"
)

func orangesRotting(grid [][]int) int {
	placeholder := 9999
	for i, _ := range grid {
		for j, _ := range grid[i] {
			switch grid[i][j] {
			case 1:
				grid[i][j] = placeholder
			case 2:
				grid[i][j] = 0
			case 0:
				grid[i][j] = -1
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				grid = BFS(grid, i, j)
				printMatrix(grid)
			}
		}
	}

	globalMax := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > globalMax {
				globalMax = grid[i][j]
			}
		}
	}

	return globalMax

}

func BFS(grid [][]int, i, j int) [][]int {

	queue := list.New()
	depth := 0
	queue.PushBack([]int{i, j, depth})

	for queue.Len() > 0 {
		oldest := queue.Front()
		queue.Remove(oldest)

		current := oldest.Value.([]int)

		ci, cj, depth := current[0], current[1], current[2]

		if ci > 0 && grid[ci-1][cj] > depth+1 {
			queue.PushBack([]int{ci - 1, cj, depth + 1})
			grid[ci-1][cj] = depth + 1
		}

		if ci < len(grid)-1 && grid[ci+1][cj] > depth+1 {
			queue.PushBack([]int{ci + 1, cj, depth + 1})
			grid[ci+1][cj] = depth + 1
		}

		if cj > 0 && grid[ci][cj-1] > depth+1 {
			queue.PushBack([]int{ci, cj - 1, depth + 1})
			grid[ci][cj-1] = depth + 1
		}

		if cj < len(grid[ci])-1 && grid[ci][cj+1] > depth+1 {
			queue.PushBack([]int{ci, cj + 1, depth + 1})
			grid[ci][cj+1] = depth + 1
		}

	}

	return grid

}

func printMatrix(matrix [][]int) {
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

func main() {
	// [[2,1,1],[1,1,0],[0,1,1]]
	orangesRotting([][]int{[]int{2, 1, 1}, []int{1, 1, 0}, []int{0, 1, 1}})
}
