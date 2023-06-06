package lc0110_2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return balancedHeight(root) != -1
}

func balancedHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := balancedHeight(root.Left), balancedHeight(root.Right)
	if l == -1 || r == -1 || abs(l-r) > 1 {
		return -1
	}
	return max(l, r) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
