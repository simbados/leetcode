package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var result []string
	var currentInterval = []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 != nums[i] {
			result = appendToResult(currentInterval, result)
			currentInterval = []int{nums[i]}
		} else {
			if len(currentInterval) <= 1 {
				currentInterval = append(currentInterval, nums[i])
			} else {
				currentInterval[1] = nums[i]
			}
		}
	}
	result = appendToResult(currentInterval, result)
	return result
}

func appendToResult(currentInterval []int, result []string) []string {
	if len(currentInterval) == 1 {
		result = append(result, fmt.Sprint(currentInterval[0]))
	} else {
		result = append(result, fmt.Sprintf("%v->%v", currentInterval[0], currentInterval[1]))
	}
	return result
}
