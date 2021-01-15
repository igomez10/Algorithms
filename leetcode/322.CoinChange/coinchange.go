package main

import "sort"

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	if amount == 0 {
		return 0
	}
	memo := map[int]int{}
	for i := range coins {
		memo[coins[i]] = 1
	}
	val := coinChangeMemo(coins, amount, &memo)

	return val
}

func coinChangeMemo(coins []int, num int, memo *map[int]int) int {
	if num < 0 {
		return -1
	}

	if val, exist := (*memo)[num]; exist {
		return val
	}

	for i := range coins {
		if num-coins[i] > 0 {
			if val := coinChangeMemo(coins, num-coins[i], memo); val > 0 {

				if oldMin, exist := (*memo)[num]; exist {
					(*memo)[num] = min(oldMin, val+1)
				} else {
					(*memo)[num] = val + 1
				}

			}
		} else {
			break
		}
	}

	if (*memo)[num] == 0 {
		(*memo)[num] = -1
	}

	return (*memo)[num]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
