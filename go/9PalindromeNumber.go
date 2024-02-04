package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	a := strconv.Itoa(x)
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-1-i] {
			return false
		}
	}
	return true
}
