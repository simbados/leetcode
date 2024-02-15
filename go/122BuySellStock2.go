package main

func maxProfit2(prices []int) int {
	margin := 0
	if len(prices) <= 1 {
		return 0
	}
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			margin += prices[i+1] - prices[i]
		}
	}
	return margin
}
