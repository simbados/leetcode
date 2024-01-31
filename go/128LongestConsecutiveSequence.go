package main

func longestConsecutive(nums []int) int {
	mapping := make(map[int]bool)
	for _, val := range nums {
		mapping[val] = true
	}
	var longest = 0
	for key := range mapping {
		if ok := mapping[key-1]; !ok {
			length := 1
			for next := mapping[key+length]; next; next = mapping[key+length] {
				length++
			}
			longest = max(length, longest)
		}
	}
	return longest
}
