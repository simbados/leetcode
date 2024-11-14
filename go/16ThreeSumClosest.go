package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	current := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[right] + nums[left]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(current-target) {
				current = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return current
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
