package main

import (
	"testing"
)

func TestReverseVowelsOfAString(t *testing.T) {
	tests := []struct {
		a        string
		expected string
	}{
        {"hello", "holle"},
        {"leetcode", "leotcede"},
        {"whatisthis", "whitisthas"},
        {"Hellomisenior", "Hollimesinoer"},
	}

	for _, tt := range tests {
		res := reverseVowels(tt.a)
        if res != tt.expected {
            t.Errorf("reverseVowels(%v) = %v; but got %v", tt.a, tt.expected, res)
        }
	}
}

