package main

// DP solution, inefficient O(n) for memory and O(n^2) runtime
func canJumpDp(nums []int) bool {
	mem := map[int]bool{}
	return helper(nums, 0, mem)
}

func helper(nums []int, index int, mem map[int]bool) bool {
	if val, exists := mem[index]; exists {
		return val
	}
	if index > len(nums)-1 {
		return false
	}
	if index == len(nums)-1 {
		return true
	}
	for i := nums[index]; i >= 1; i-- {
		val := helper(nums, i+index, mem)
		mem[index+i] = val
		if val {
			return true
		}
	}
	return false
}

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	necessary := 1
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] >= necessary {
			necessary = 1
			continue
		}
		necessary++
	}
	return nums[0] >= necessary
}
