package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + ((end - start) / 2)
		if mid-1 < 0 && nums[mid] > nums[mid+1] {
			return mid
		} else if mid+1 > len(nums)-1 && nums[mid] > nums[mid-1] {
			return mid
		} else if nums[mid+1] > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
