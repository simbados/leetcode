package main

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	tests := []struct {
		a        int
		expected bool
	}{
		{121, true},
		{-121, false},
		{1112222, false},
		{112211, true},
	}

	for _, tt := range tests {
		res := isPalindrome(tt.a)
		if res != tt.expected {
			t.Errorf("isPalindrom(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
