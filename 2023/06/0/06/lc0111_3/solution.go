package lc0111

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := minDepth(root.Left), minDepth(root.Right)
	if l == 0 || r == 0 {
		return l + r + 1
	}
	return min(l, r) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
