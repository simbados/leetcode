package main

import (
	"testing"
)

func TestFirstUniqueCharacter(t *testing.T) {
	tests := []struct {
		a        string
		expected int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
		{"z", 0},
		{"dddccdbba", 8},
	}

	for _, tt := range tests {
		res := firstUniqChar(tt.a)
		if res != tt.expected {
			t.Errorf("firstUniqChar(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
