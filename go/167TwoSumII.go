package main

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		start := i + 1
		end := len(numbers) - 1
		sum := target - numbers[i]
		for start <= end {
			mid := start + ((end - start) / 2)
			if mid > end {
				mid = end
			}
			if numbers[mid] == sum {
				return []int{i + 1, mid + 1}
			}
			if numbers[mid] < sum {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return []int{}
}
