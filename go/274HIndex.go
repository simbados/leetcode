package main

// [1, 1, 100, 100, 100]
func hIndex(citations []int) int {
	lenC := len(citations)
	count := make([]int, lenC+1)
	for _, v := range citations {
		if lenC <= v {
			count[lenC] += 1
		} else {
			count[v] += 1
		}
	}
	var counting int
	for i := lenC; i >= 0; i-- {
		counting += count[i]
		if counting >= i {
			return i
		}
	}
	return 0
}
