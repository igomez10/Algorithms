package main

func minDifficulty(jobs []int, d int) int {
	if len(jobs) < d {
		return -1
	}
	memo := make([][]int, len(jobs))
	for i := range memo {
		memo[i] = make([]int, d)
	}
	val := numPartitions(jobs, d-1, 0, &memo)

	return val
}

func numPartitions(jobs []int, partitionsLeft, start int, memo *[][]int) int {
	if start >= len(jobs) {
		return -1
	}

	if val := (*memo)[start][partitionsLeft]; val > 0 {
		return val
	}

	if partitionsLeft == 0 {
		mostDifficult := max(jobs[start:]...)
		(*memo)[start][partitionsLeft] = mostDifficult
		return mostDifficult
	}

	minimum := 0
	for i := start + 1; i < len(jobs)-partitionsLeft+1; i++ {
		current := max(jobs[start:i]...) + numPartitions(jobs, partitionsLeft-1, i, memo)
		if minimum == 0 {
			minimum = current
		} else {
			minimum = min(minimum, current)
		}
	}

	(*memo)[start][partitionsLeft] = minimum

	return minimum
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
