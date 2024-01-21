package main

import (
	"testing"
)

func TestIsomorphicString(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected bool
	}{
		{"egg", "add", true},
		{"eggs", "add", false},
		{"what", "rats", true},
		{"badc", "baba", false},
	}

	for _, tt := range tests {
		res := isIsomorphic(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("isIsomorphic(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
