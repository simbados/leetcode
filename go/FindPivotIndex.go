package main

func pivotIndex(nums []int) int {
	left := 0
	sum := sum(nums)
	for index, val := range nums {
		sum = sum - val
		if left == sum {
			return index
		}
		left = left + val
	}
	return -1
}

func sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
