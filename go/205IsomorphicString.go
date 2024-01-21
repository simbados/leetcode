package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	helper1 := make(map[rune]rune)
	helper2 := make(map[rune]rune)
	second := []rune(t)
	for index, value := range s {
		val, ok := helper1[value]
		if ok && val != second[index] {
			return false
		} else {
			helper1[value] = second[index]
		}
		val2, ok2 := helper2[second[index]]
		if ok2 && val2 != value {
			return false
		} else {
			helper2[second[index]] = value
		}
	}
	return true
}
