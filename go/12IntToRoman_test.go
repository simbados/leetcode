package main

import (
	"testing"
)

func TestIntegerToRoman(t *testing.T) {
	tests := []struct {
		a        int
		expected string
	}{
		{3, "III"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, tt := range tests {
		res := intToRoman(tt.a)
		if res != tt.expected {
			t.Errorf("intToRoman(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
