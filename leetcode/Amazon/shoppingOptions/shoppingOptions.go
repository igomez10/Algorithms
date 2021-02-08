package main

import "sort"

func getShoppingOptions(jeans, shoes, skirts, tops []int, budget int) int {

	sort.Ints(jeans)
	sort.Ints(shoes)
	sort.Ints(skirts)
	sort.Ints(tops)
	counterNumWays := 0
	for i := 0; i < len(jeans) && jeans[i] < budget; i++ {
		oneItem := jeans[i]
		for j := 0; j < len(shoes) && (oneItem+shoes[j]) < budget; j++ {
			twoItem := oneItem + shoes[j]
			for z := 0; z < len(skirts) && twoItem+skirts[z] < budget; z++ {
				threeItem := twoItem + skirts[z]
				remaining := budget - threeItem
				// find biggest value below `remaining`
				idxBiggest := binarySearch(tops, remaining+1)
				counterNumWays += idxBiggest + 1
			}
		}
	}
	return counterNumWays
}

func binarySearch(arr []int, target int) int {
	lo := 0
	hi := len(arr) - 1
	ans := -1
	for lo <= hi {
		middle := (lo + hi) / 2
		if target < arr[middle] {
			hi = middle - 1
		} else {
			ans = middle
			lo = middle + 1
		}
	}
	return ans
}
