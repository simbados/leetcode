package main

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		a        uint32
		expected uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}

	for _, tt := range tests {
		res := reverseBits(tt.a)
		if res != tt.expected {
			t.Errorf("reverseBits(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
