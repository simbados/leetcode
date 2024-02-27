package main

import (
	"testing"
)

func TestReverseWordString(t *testing.T) {
	tests := []struct {
		a        string
		expected string
	}{
		{" Hello World ", "World Hello"},
		{"the sky is blue", "blue is sky the"},
		{"a good   example", "example good a"},
	}

	for _, tt := range tests {
		res := reverseWords(tt.a)
		if res != tt.expected {
			t.Errorf("reverseWords(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
