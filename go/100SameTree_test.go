package main

import (
	"testing"
)

func TestOne(t *testing.T) {
	treeNode2 := TreeNode{Left: nil, Right: nil, Val: 2}
	treeNode3 := TreeNode{Left: nil, Right: nil, Val: 3}
	treeNode1 := TreeNode{Left: &treeNode2, Right: &treeNode3, Val: 1}
	treeNodeCopy2 := TreeNode{Left: nil, Right: nil, Val: 2}
	treeNodeCopy3 := TreeNode{Left: nil, Right: nil, Val: 3}
	treeNodeCopy1 := TreeNode{Left: &treeNodeCopy2, Right: &treeNodeCopy3, Val: 1}
	tests := []struct {
		a        *TreeNode
		b        *TreeNode
		expected bool
	}{
		{&treeNode1, &treeNodeCopy1, true},
	}

	for _, tt := range tests {
		res := isSameTree(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("isSameTree(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}

func TestTwo(t *testing.T) {
	treeNode2 := TreeNode{Left: nil, Right: nil, Val: 2}
	treeNode1 := TreeNode{Left: nil, Right: &treeNode2, Val: 1}
	treeNodeCopy2 := TreeNode{Left: nil, Right: nil, Val: 2}
	treeNodeCopy1 := TreeNode{Left: &treeNodeCopy2, Right: nil, Val: 1}
	tests := []struct {
		a        *TreeNode
		b        *TreeNode
		expected bool
	}{
		{&treeNode1, &treeNodeCopy1, false},
	}

	for _, tt := range tests {
		res := isSameTree(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("isSameTree(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}

func TestThree(t *testing.T) {
	treeNode2 := TreeNode{Left: nil, Right: nil, Val: 1}
	treeNode3 := TreeNode{Left: nil, Right: nil, Val: 2}
	treeNode1 := TreeNode{Left: &treeNode3, Right: &treeNode2, Val: 1}
	treeNodeCopy2 := TreeNode{Left: nil, Right: nil, Val: 1}
	treeNodeCopy3 := TreeNode{Left: nil, Right: nil, Val: 2}
	treeNodeCopy1 := TreeNode{Left: &treeNodeCopy2, Right: &treeNodeCopy3, Val: 1}
	tests := []struct {
		a        *TreeNode
		b        *TreeNode
		expected bool
	}{
		{&treeNode1, &treeNodeCopy1, false},
	}

	for _, tt := range tests {
		res := isSameTree(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("isSameTree(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
