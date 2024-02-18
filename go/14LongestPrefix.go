package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	a := strs[0]
	for i := 1; i < len(strs); i++ {
		b := strs[i]
		if a == "" || b == "" {
			return ""
		}
		if len(a) > len(b) {
			c := a
			a = b
			b = c
		}
		for len(a) > 0 {
			if strings.HasPrefix(b, a) {
				break
			}
			a = a[0 : len(a)-1]
		}
	}
	return a
}
