package main

// O(n^2)
func twoSumReal(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// O(n)
func twoSumOne(numbers []int, target int) []int {
	mapping := make(map[int]int)
	for i, val := range numbers {
		remainder := target - val
		if index, ok := mapping[remainder]; ok {
			return []int{index, i}
		}
		mapping[val] = i
	}
	return []int{}
}
