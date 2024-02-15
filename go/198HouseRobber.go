package main

// Iterative
func rob1(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i <= len(nums)-1; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

// Recursive
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	var dpF func(i int) int
	dpF = func(i int) int {
		if i == 0 {
			return nums[0]
		} else if i == 1 {
			return max(nums[i], nums[0])
		}
		dp[i] = max(dpF(i-2)+nums[i], dpF(i-1))
		return dp[i]
	}
	return dpF(len(nums) - 1)
}
