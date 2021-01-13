package main

func wordBreak(s string, dict []string) bool {
	dp := make([]bool, len(s)+1)
	dictmap := map[string]bool{}
	for i := range dict {
		dictmap[dict[i]] = true
	}

	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dictmap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}
