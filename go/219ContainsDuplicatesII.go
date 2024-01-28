package main

import (
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	mapping := make(map[int]int)
	minVal := math.MaxInt32
	for index, val := range nums {
		if previousIndex, ok := mapping[val]; ok {
			minVal = index - previousIndex
			if minVal <= k {
				return true
			}
		}
		mapping[val] = index
	}
	return false
}
