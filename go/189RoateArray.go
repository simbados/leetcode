package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	buff := make([]int, 0, len(nums)-1)
	buff = append(buff, nums[len(nums)-k:]...)
	buff = append(buff, nums[:len(nums)-k]...)
	copy(nums, buff)
}
