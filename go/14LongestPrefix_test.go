package main

import (
	"testing"
)

func TestLongestPrefix(t *testing.T) {
	tests := []struct {
		a        []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, tt := range tests {
		res := longestCommonPrefix(tt.a)
		if res != tt.expected {
			t.Errorf("longestPrefix(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
