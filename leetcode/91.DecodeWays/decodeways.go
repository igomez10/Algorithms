package main

func numDecodings(s string) int {
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
