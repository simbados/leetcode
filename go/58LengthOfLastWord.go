package main

import "strings"

func lengthOfLastWord(s string) int {
	splitArr := strings.Split(strings.Trim(s, " "), " ")
	return len(splitArr[len(splitArr)-1])
}
