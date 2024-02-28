package main

import (
	"testing"
)

// "HelloThisIsMario" 5
// H   S
// E  I I     O
// L H   S   I
// LT     M R
// O       A
func TestZigZag(t *testing.T) {
	tests := []struct {
		a        string
		b        int
		expected string
	}{
    {"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
    {"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
    {"HELLOTHISISMARIO", 5, "HSEIIOLHSILTMROA"},
	}

	for _, tt := range tests {
		res := convert(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("convert(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
