package lc0124

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var sideMax func(root *TreeNode) int
	sideMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := sideMax(root.Left), sideMax(root.Right)
		ans = max(ans, left+root.Val+right)
		return max(0, max(left, right)+root.Val)
	}
	sideMax(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
