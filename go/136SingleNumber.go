package main

func singleNumber(nums []int) int {
	start := 0
	for _, val := range nums {
		start = start ^ val
	}
	return start
}
