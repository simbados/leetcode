package main

// Complicated
func rangeBitwiseAnd2(left int, right int) int {
	if left == right {
		return left
	}
	if left == 0 || right == 0 {
		return 0
	}
	leftOuterBit, rightOuterBit := 0, 0
	for i := 0; i < 32; i++ {
		if (left>>i)&1 == 1 {
			leftOuterBit = i
		}
		if (right>>i)&1 == 1 {
			rightOuterBit = i
		}
	}
	if rightOuterBit == leftOuterBit {
		result := left
		for i := left + 1; i <= right; i++ {
			result = result & i
		}
		return result
	}
	return 0
}

// Easy
func rangeBitwiseAnd(left int, right int) int {
	count := 0
	for left != right {
		left >>= 1
		right >>= 1
		count++
	}
	return left << count
}
