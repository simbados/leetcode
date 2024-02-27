package main

import (
	"strings"
)

func reverseWords(s string) string {
	splitArr := strings.Split(strings.Trim(s, " "), " ")
    result := ""
    for i := len(splitArr) - 1; i >= 0; i-- {
        if splitArr[i] != "" {
            result += splitArr[i] + " "
        }
    }
    return strings.Trim(result, " ")
}
