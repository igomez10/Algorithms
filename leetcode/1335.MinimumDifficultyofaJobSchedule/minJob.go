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

// 6 5 4 3 2 1 , d = 2
// 6;    5 4 3 2 1 , d=1
// 6;5;  4 3 2 1 , d=0
func minDifficultyMemo(jobs []int, daysLeft, start int, memo [][]int) int {
	if daysLeft == 0 {
		if start == len(jobs) {
			return 0 // reached the end of the array, no jobs remaning so return 0
		}

		return math.MaxInt32 // reached the end of days and has jobs to do, infinite difficulty
	}

	if memo[daysLeft][start] != -1 {
		return memo[daysLeft][start] // solved this problem before
	}

	currentMax := jobs[start]
	currentMin := math.MaxInt32
	for i := start; i <= len(jobs)-daysLeft; i++ { // start at every possible day from the previous start and the jobs-daysleft start
		currentMax = max(currentMax, jobs[i])
		newMin := currentMax + minDifficultyMemo(jobs, daysLeft-1, i+1, memo) // check every possible difficulty from every possible starting point
		currentMin = min(currentMin, newMin)
	}

	memo[daysLeft][start] = currentMin
	return memo[daysLeft][start]
}

func minDifficultyBottomUP(jobs []int, days int) int {
	if days > len(jobs) {
		return -1
	}

	dp := make([][]int, days)
	for i := range dp {
		dp[i] = make([]int, len(jobs))
	}

	dp[0][0] = jobs[0]
	for i := 1; i < len(jobs); i++ {
		dp[0][i] = max(jobs[i], dp[0][i-1]) // keep highest seen
	}

	for currentDay := 1; currentDay < days; currentDay++ {

		for currentJob := currentDay; currentJob < len(jobs); currentJob++ {
			localMax := jobs[currentJob]
			dp[currentDay][currentJob] = math.MaxInt32

			for r := currentJob; r >= currentDay; r-- {
				localMax = max(localMax, jobs[r])
				dp[currentDay][currentJob] = min(dp[currentDay][currentJob], dp[currentDay-1][r-1]+localMax)
			}
		}
	}

	return dp[days-1][len(jobs)-1]
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
