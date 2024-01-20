package main

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	answer := math.MaxInt32
	left := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			answer = min(answer, i+1-left)
			sum -= nums[left]
			left++
		}
	}
	if answer == math.MaxInt32 {
		return 0
	} else {
		return answer
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
