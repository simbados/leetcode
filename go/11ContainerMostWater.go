package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	answer := 0
	for left <= right {
		answer = max(area(left, right, height), answer)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return answer
}

func area(i1 int, i2 int, height []int) int {
	return (i2 - i1) * min(height[i1], height[i2])
}
