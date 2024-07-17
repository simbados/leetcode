package main

import (
	"testing"
)

func TestMergeStringsAlternately(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
	}

	for _, tt := range tests {
		res := mergeStringsAlternately(tt.a, tt.b)
        if res != tt.expected {
            t.Errorf("mergeStringsAlternately(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
        }
	}
}

