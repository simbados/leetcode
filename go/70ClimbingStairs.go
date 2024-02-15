package main

func climbStairs(n int) int {
	mem := map[int]int{}
	var dp func(left int) int
	dp = func(left int) int {
		if left < 0 {
			return 0
		}
		if left <= 1 {
			return 1
		}
		if val, exists := mem[left]; exists {
			return val
		}
		mem[left-1] = dp(left - 1)
		mem[left-2] = dp(left - 2)
		return mem[left-1] + mem[left-2]
	}
	return dp(n)
}
