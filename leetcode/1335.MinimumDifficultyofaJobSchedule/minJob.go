package main

import (
	"math"
)

func minDifficulty(jobs []int, d int) int {
	if len(jobs) < d {
		return -1
	}
	memo := make([][]int, d+1)
	for i := range memo {
		memo[i] = make([]int, len(jobs)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	val := minDifficultyMemo(jobs, d, 0, memo)

	return val
}

func minDifficultyMemo(jobs []int, daysLeft, start int, memo [][]int) int {
	if daysLeft == 0 && start == len(jobs) {
		return 0
	}

	if daysLeft == 0 {
		return math.MaxInt32
	}

	if memo[daysLeft][start] != -1 {
		return memo[daysLeft][start]
	}

	currentMax := jobs[start]
	currentMin := math.MaxInt32
	for i := start; i <= len(jobs)-daysLeft; i++ {
		currentMax = max(currentMax, jobs[i])
		newMin := currentMax + minDifficultyMemo(jobs, daysLeft-1, i+1, memo)
		currentMin = min(currentMin, newMin)
	}
	memo[daysLeft][start] = currentMin
	return memo[daysLeft][start]
}

func min(args ...int) int {
	globalMin := args[0]
	for i := range args {
		if args[i] < globalMin {
			globalMin = args[i]
		}
	}
	return globalMin
}

func max(args ...int) int {
	globalMax := args[0]
	for i := range args {
		if args[i] > globalMax {
			globalMax = args[i]
		}
	}
	return globalMax
}
