package main

func isSubsequence(s string, t string) bool {
	i := 0
	runes := []rune(s)
	for _, value := range t {
		if i == len(s) {
			break
		}
		if value == runes[i] {
			i++
		}
	}
	if i == len(s) {
		return true
	} else {
		return false
	}
}
