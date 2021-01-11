package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	cheapest := prices[0]
	profit := 0
	for i := range prices {

		if prices[i]-cheapest > profit {
			profit = prices[i] - cheapest
		}

		if prices[i] < cheapest {
			cheapest = prices[i]
		}
	}

	return profit
}
