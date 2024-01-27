package main

import "math"

func maxProfit(prices []int) int {
	profit := 0
	minV := math.MaxInt32
	for _, val := range prices {
		if val < minV {
			minV = val
		} else if tmp := val - minV; tmp > profit {
			profit = tmp
		}
	}
	return profit
}
