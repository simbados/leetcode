package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	mappinger := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var result []string
	var backtracking func(digits string, combis string)
	backtracking = func(digits string, combis string) {
		if digits == "" {
			result = append(result, combis)
		} else {
			current := mappinger[digits[0]-'2']
			for _, char := range current {
				backtracking(digits[1:], combis+string(char))
			}
		}
	}
	backtracking(digits, "")
	return result
}
