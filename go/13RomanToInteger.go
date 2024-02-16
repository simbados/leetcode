package main

var mapping = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

// MCMXCIV
func romanToInt(s string) int {
	sum := 0
	i := 0
	for i < len(s) {
		current := mapping[rune(s[i])]
		if i == len(s)-1 {
			sum += current
			break
		}
		next := mapping[rune(s[i+1])]
		if next > current {
			sum += next - current
			i = i + 2
		} else {
			sum += current
			i = i + 1
		}
	}
	return sum
}
