package main

import (
	"math"
)

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	mem := make([]int, len(nums))
	mem[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		mem[i] = math.MaxInt32
		for j := nums[i]; j > 0; j-- {
			if i+j < len(nums)-1 {
				mem[i] = min(mem[i], mem[i+j]+1)
			} else {
				mem[i] = 1
				break
			}
		}
	}
	return mem[0]
}
