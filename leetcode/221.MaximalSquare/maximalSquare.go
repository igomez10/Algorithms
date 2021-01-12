package main

func maximalSquare(arr [][]byte) int {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return 0
	}

	dp := make([][]byte, len(arr))
	var maxWidth byte = 0
	for i := range dp {
		dp[i] = make([]byte, len(arr[0]))
	}

	for i := range arr {
		for j := range arr[i] {
			var closeMin byte = 0
			if arr[i][j] == 1 {
				if j > 0 && i > 0 {
					closeMin = dp[i][j-1]
					closeMin = min(closeMin, dp[i-1][j])
					closeMin = min(closeMin, dp[i-1][j-1])
				}
				dp[i][j] = closeMin + 1
				if dp[i][j] > maxWidth {
					maxWidth = dp[i][j]
				}
			}
		}
	}

	return int(maxWidth)
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}
