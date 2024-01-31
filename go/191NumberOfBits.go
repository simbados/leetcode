package main

func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		count += int(num % 2)
		num = num >> 1
	}
	return count
}
