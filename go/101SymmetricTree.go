package main

func isSymmetric(root *TreeNode) bool {
	return sameValue(root.Left, root.Right)
}

func sameValue(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	if right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return sameValue(left.Left, right.Right) && sameValue(left.Right, right.Left)
}
