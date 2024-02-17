package main

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		a        string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, tt := range tests {
		res := lengthOfLastWord(tt.a)
		if res != tt.expected {
			t.Errorf("lengthOfLastWord(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
