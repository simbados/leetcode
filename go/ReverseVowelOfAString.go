package main

import "strings"

func reverseVowels(s string) string {
    start, end := 0, len(s) - 1
    stringAsRune := make([]rune, len(s))
    for i, r := range s {
        stringAsRune[i] = r
    }
    for start < end {
        if !isVowel(stringAsRune[start]) {
            start++
            continue
        } 
        if !isVowel(stringAsRune[end]) {
            end--
            continue
        }
        stringAsRune[start], stringAsRune[end] = stringAsRune[end], stringAsRune[start]
        start++
        end--
    }
    return string(stringAsRune)
}

func isVowel(s rune) bool {
    toCheck := strings.ToLower(string(s)) 
    if toCheck == "a" || toCheck == "e" || toCheck == "i" || toCheck == "o" || toCheck == "u" {
        return true
    }
    return false
}
