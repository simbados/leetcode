package _go

import (
	"math"
	"sort"
)

func removeElement(nums []int, val int) int {
	var count = 0
	for i, numsVal := range nums {
		if numsVal == val {
			nums[i] = math.MaxInt32
		} else {
			count++
		}
	}
	sort.Ints(nums)
	return count
}
