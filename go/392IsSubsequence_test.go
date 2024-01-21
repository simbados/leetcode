package main

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"abb", "abb", true},
		{"", "", true},
		{"", "asd", true},
	}

	for _, tt := range tests {
		result := isSubsequence(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("isSubsequence(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, result)
		}
	}
}
