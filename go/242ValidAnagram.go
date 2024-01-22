package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	helper := make(map[rune]rune)
	for _, val := range s {
		if existing, ok := helper[val]; ok {
			helper[val] = existing + 1
		} else {
			helper[val] = 1
		}
	}
	for _, val := range t {
		if existing, ok := helper[val]; ok {
			helper[val] = existing - 1
		} else {
			return false
		}
	}
	for _, value := range helper {
		if value != 0 {
			return false
		}
	}

	return true
}
