package main

type TreeNode struct {
	Val   int       `json:"Val"`
	Left  *TreeNode `json:"Left"`
	Right *TreeNode `json:"Right"`
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return findVal(p, q)
}

func findVal(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return findVal(p.Left, q.Left) && findVal(p.Right, q.Right)
}
