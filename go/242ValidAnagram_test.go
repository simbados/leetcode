package main

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, tt := range tests {
		res := isAnagram(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("isAnagram(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
