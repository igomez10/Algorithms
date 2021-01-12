package main

import (
	"strconv"
)

func numDecodings(s string) int {
	memo := map[int]int{}
	return recursiveWithMemo(s, 0, &memo)
}

func recursiveWithMemo(s string, index int, memo *map[int]int) int {
	if index == len(s) {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	if index == len(s)-1 {
		return 1
	}

	if ans, exist := (*memo)[index]; exist {
		return ans
	}

	ans := recursiveWithMemo(s, index+1, memo)
	if num, err := strconv.Atoi(string(s[index : index+2])); err == nil && num <= 26 {
		ans += recursiveWithMemo(s, index+2, memo)
	}
	(*memo)[index] = ans
	return ans
}

func numDecodingsIterative(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	if s[0] == '0' {
		dp[1] = 0
	}

	for i := 2; i < len(dp); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		if s[i-2] == '1' || (s[i-2] == '2' &&
			(s[i-1] == '0' ||
				s[i-1] == '1' ||
				s[i-1] == '2' ||
				s[i-1] == '3' ||
				s[i-1] == '4' ||
				s[i-1] == '5' ||
				s[i-1] == '6')) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(dp)-1]
}
