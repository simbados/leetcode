package main

import (
	"slices"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		a        []byte
		expected []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, tt := range tests {
		b := slices.Clone(tt.a)
		reverseString(tt.a)
		if !slices.Equal(tt.a, tt.expected) {
			t.Errorf("reverseString(%v) = %v; but got %v", b, tt.a, tt.expected)
		}
	}
}
