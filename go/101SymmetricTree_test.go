package main

import "testing"

func TestSymmetricTree(t *testing.T) {
	treeNode7 := TreeNode{Left: nil, Right: nil, Val: 4}
	treeNode6 := TreeNode{Left: nil, Right: nil, Val: 4}
	treeNode5 := TreeNode{Left: nil, Right: nil, Val: 3}
	treeNode4 := TreeNode{Left: nil, Right: nil, Val: 3}
	treeNode2 := TreeNode{Left: &treeNode5, Right: &treeNode6, Val: 2}
	treeNode3 := TreeNode{Left: &treeNode7, Right: &treeNode4, Val: 2}
	treeNode1 := TreeNode{Left: &treeNode2, Right: &treeNode3, Val: 1}
	tests := []struct {
		a        *TreeNode
		expected bool
	}{
		{&treeNode1, true},
	}

	for _, tt := range tests {
		res := isSymmetric(tt.a)
		if res != tt.expected {
			t.Errorf("isSymmetric(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
