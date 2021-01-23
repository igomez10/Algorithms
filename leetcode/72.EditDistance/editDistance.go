package main

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	if word1 == "" || word2 == "" {
		return len(word2) + len(word1)
	}

	for i := range dp {
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				ans1 := dp[i-1][j] + 1
				ans2 := dp[i][j-1] + 1
				ans3 := dp[i-1][j-1] + 1
				ans := min(ans1, ans2, ans3)
				dp[i][j] = ans
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(args ...int) int {
	smallest := args[0]
	for i := range args {
		if args[i] < smallest {
			smallest = args[i]
		}
	}
	return smallest
}

// if word1[0] == word2[0] {
// 	ans := minDistanceMemo(word1[1:], word2[1:], dp)
// 	dp[len(word1)][len(word2)] = ans
// 	return ans
// }

// ans1 := 1 + minDistanceMemo(word1[1:], word2, dp)
// ans2 := 1 + minDistanceMemo(word1, word2[1:], dp)
// ans3 := 1 + minDistanceMemo(word1[1:], word2[1:], dp)

// ans := min(ans1, ans2, ans3)
// dp[len(word1)][len(word2)] = ans
