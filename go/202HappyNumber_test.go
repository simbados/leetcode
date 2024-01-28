package main

import (
	"testing"
)

func TestHappyNumber(t *testing.T) {
	tests := []struct {
		a        int
		expected bool
	}{
		{19, true},
		{2, false},
	}

	for _, tt := range tests {
		res := isHappy(tt.a)
		if res != tt.expected {
			t.Errorf("isHappy(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
