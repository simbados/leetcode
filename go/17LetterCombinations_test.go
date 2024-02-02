package main

import (
	"slices"
	"testing"
)

func TestLetterCombination(t *testing.T) {
	tests := []struct {
		a        string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		res := letterCombinations(tt.a)
		if !slices.Equal(res, tt.expected) {
			t.Errorf("letterCombination(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
