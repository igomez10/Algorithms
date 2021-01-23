package main

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1)+1)
	for i := range memo {
		memo[i] = make([]int, len(word2)+1)
	}

	return minDistanceMemo(word1, word2, memo)
}

func minDistanceMemo(word1, word2 string, memo [][]int) int {
	if word1 == "" || word2 == "" {
		return len(word2) + len(word1)
	}

	if memo[len(word1)][len(word2)] > 0 {
		return memo[len(word1)][len(word2)]
	}

	if word1[0] == word2[0] {
		ans := minDistanceMemo(word1[1:], word2[1:], memo)
		memo[len(word1)][len(word2)] = ans
		return ans
	}

	ans1 := 1 + minDistanceMemo(word1[1:], word2, memo)
	ans2 := 1 + minDistanceMemo(word1, word2[1:], memo)
	ans3 := 1 + minDistanceMemo(word1[1:], word2[1:], memo)

	ans := min(ans1, ans2, ans3)
	memo[len(word1)][len(word2)] = ans

	return memo[len(word1)][len(word2)]
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
