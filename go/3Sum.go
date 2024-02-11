package main

import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	mapping := map[string]bool{}
	var solution [][]int
	if len(nums) < 3 {
		return solution
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[right] + nums[left]
			if sum == 0 {
				value := []int{nums[i], nums[left], nums[right]}
				sortedCopy := make([]int, len(value))
				copy(sortedCopy, value)
				sort.Ints(sortedCopy)
				hashValue := strconv.Itoa(sortedCopy[0]) + strconv.Itoa(sortedCopy[1]) + strconv.Itoa(sortedCopy[2])
				if exists := mapping[hashValue]; !exists {
					solution = append(solution, value)
					mapping[hashValue] = true
				}
				right = right - 1
			} else {
				if sum > 0 {
					right--
				} else {
					left++
				}
			}
		}
	}
	return solution
}
