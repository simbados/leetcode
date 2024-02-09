package main

func permute(nums []int) [][]int {
	var result [][]int
	var backtracking func(values []int, accumulated []int)
	backtracking = func(index []int, accumulated []int) {
		if len(index) == 0 {
			result = append(result, accumulated)
			return
		}
		for i, num := range index {
			backtracking(filterIndex(index, i), append(accumulated, num))
		}
	}
	backtracking(nums, []int{})
	return result
}

func filterIndex(values []int, index int) []int {
	copyValues := make([]int, len(values))
	copy(copyValues, values)
	if index == len(copyValues)-1 {
		return copyValues[:index]
	} else if index == 0 {
		return copyValues[1:]
	}
	return append(copyValues[:index], copyValues[index+1:]...)
}
