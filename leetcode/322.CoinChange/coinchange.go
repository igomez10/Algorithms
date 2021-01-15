package main

import "math"

func coinChangeBottomUp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := range dp {
		for j := range coins {
			if i-coins[j] >= 0 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChangeTopBottom(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	dp := make([]int, amount)

	return coinChangeMemo(coins, amount, &dp)

}

func coinChangeMemo(coins []int, amount int, dp *[]int) int {
	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	if val := (*dp)[amount-1]; val != 0 {
		return val
	}

	minAmount := math.MaxInt16
	for i := range coins {
		if amount-coins[i] >= 0 {
			res := coinChangeMemo(coins, amount-coins[i], dp)
			if res >= 0 && res < minAmount {
				minAmount = res + 1

			}

		}
	}
	if minAmount != math.MaxInt16 {
		(*dp)[amount-1] = minAmount
	} else {
		(*dp)[amount-1] = -1
	}

	return (*dp)[amount-1]
}
