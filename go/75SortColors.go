package main

// Bubble Sort
func bubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				buffer := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = buffer
			}
		}
	}
}
