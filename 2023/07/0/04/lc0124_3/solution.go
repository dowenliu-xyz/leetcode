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
	// 注意：有负数节点，对路径和最大化有反作用
	ans := math.MinInt // 路径和有可能是负数，所以不能用 0 做初始值
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := max(0, helper(root.Left)), max(0, helper(root.Right)) // 左右子节点的结果都有可能是负数，要去除负作用
		ans = max(ans, l+r+root.Val)                                  // 更新结果
		return max(l, r) + root.Val
	}
	helper(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
