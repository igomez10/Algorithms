package main

func DeletionDistance(source, target string) int {
	if source == "" || target == "" {
		return len(source) + len(target)
	}

	dp := make([][]int, len(source)+1)
	for i := range dp {
		dp[i] = make([]int, len(target)+1)
	}

	for i := range dp {
		dp[i][0] = i
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if source[i-1] == target[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
