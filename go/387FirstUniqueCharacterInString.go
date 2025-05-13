package main

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}
	var curr uint8
	found := make(map[uint8]bool)
	for i := 0; i < len(s)-1; i++ {
		if found[s[i]] {
			continue
		}
		curr = s[i]
		for j := i + 1; j < len(s); j++ {
			if curr == s[j] {
				found[s[j]] = true
				break
			}
			if j == len(s)-1 {
				return i
			}
		}
	}
	if !found[s[len(s)-1]] {
		return len(s) - 1
	}
	return -1
}
