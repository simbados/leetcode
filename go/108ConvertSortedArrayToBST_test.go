package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConvertSortedArray(t *testing.T) {
	treeNode2 := TreeNode{Left: nil, Right: nil, Val: -10}
	treeNode3 := TreeNode{Left: &treeNode2, Right: nil, Val: -3}
	treeNode4 := TreeNode{Left: nil, Right: nil, Val: 5}
	treeNode5 := TreeNode{Left: &treeNode4, Right: nil, Val: 9}
	treeNode1 := TreeNode{Left: &treeNode3, Right: &treeNode5, Val: 0}
	tests := []struct {
		a        []int
		expected *TreeNode
	}{
		{[]int{-10, -3, 0, 5, 9}, &treeNode1},
	}

	for _, tt := range tests {
		res := sortedArrayToBST(tt.a)
		same := isSameTree(res, tt.expected)
		if !same {
			prettyJson, _ := json.MarshalIndent(tt.expected, "", "    ")
			fmt.Println(string(prettyJson))
			prettyJson2, _ := json.MarshalIndent(res, "", "    ")
			fmt.Println(string(prettyJson2))
			t.Errorf("sortedArrayToBST(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
