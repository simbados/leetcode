package main

import (
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	tests := []struct {
		a        string
		expected int
	}{
		{"III", 3},
		{"XI", 11},
		{"MCMXCIV", 1994},
	}

	for _, tt := range tests {
		res := romanToInt(tt.a)
		if res != tt.expected {
			t.Errorf("romanToInt(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
