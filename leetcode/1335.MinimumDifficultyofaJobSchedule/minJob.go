package main

import "math"

func minDifficulty(jobs []int, d int) int {
	if len(jobs) < d {
		return -1
	}
	memo := make([][]int, d)
	for i := range memo {
		memo[i] = make([]int, len(jobs))
	}
	val := minDifficultyMemo(jobs, d-1, 0, memo)

	return val
}

func minDifficultyMemo(jobs []int, daysLeft, start int, memo [][]int) int {
	if memo[daysLeft][start] > 0 {
		return memo[daysLeft][start]
	}

	if daysLeft == 0 {
		biggest := max(jobs[start:]...)
		memo[daysLeft][start] = biggest
		return memo[daysLeft][start]
	}

	minimum := math.MaxInt32
	currentMax := jobs[start]
	for i := start + 1; i <= len(jobs)-daysLeft; i++ {
		if jobs[i-1] > currentMax {
			currentMax = jobs[i-1]
		}

		current := currentMax + minDifficultyMemo(jobs, daysLeft-1, i, memo)
		minimum = min(minimum, current)
	}

	memo[daysLeft][start] = minimum

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
