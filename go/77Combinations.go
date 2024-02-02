package main

import (
	"slices"
)

func combine(n int, k int) [][]int {
	var result [][]int
	var backtracking func(start int, combination []int)
	backtracking = func(start int, combination []int) {
		if len(combination) == k {
			result = append(result, combination)
		} else {
			for i := start; i <= n; i++ {
				backtracking(i+1, slices.Clone(append(combination, i)))
			}
		}
	}
	backtracking(1, []int{})
	return result
}
