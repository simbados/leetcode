package main

import (
	"testing"
)

func TestGreatestCommonDivisorOfStrings(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
        {"ABC", "ABCABC", "ABC"},
        {"BC", "BCBC", "BC"},
        {"LEET", "CODE", ""},
	}

	for _, tt := range tests {
		res := gcdOfStrings(tt.a, tt.b)
        if res != tt.expected {
            t.Errorf("greatestCommonDivisorOfStrings(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
        }
	}
}

