package main

func plusOne(digits []int) []int {
	over := 1
	for i := len(digits) - 1; i >= 0; i-- {
		current := digits[i]
		digits[i] = (over + current) % 10
		over = (current + over) / 10
	}
	if over == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
