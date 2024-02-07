package main

// O(n)
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	previous := 1
	for i := 2; i <= x/2; i++ {
		if i*i > x {
			break
		}
		previous = i
	}
	return previous
}

func binarySqrt(x int) int {
	left, right := 0, x+1
	for left < right {
		mid := left + ((right - left) / 2)
		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
