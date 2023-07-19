package lc0098

import (
	"math"

	. "github.com/dowenliu-xyz/leetcode/internal/types"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) bool {
	var dfs func(root *TreeNode, gt, lt int) bool
	dfs = func(root *TreeNode, gt, lt int) bool {
		if root == nil {
			return true
		}
		if root.Val <= gt || root.Val >= lt {
			return false
		}
		return dfs(root.Left, gt, root.Val) && dfs(root.Right, root.Val, lt)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

func solution2(root *TreeNode) bool {
	var inorder []int
	var visit func(root *TreeNode)
	visit = func(root *TreeNode) {
		if root == nil {
			return
		}
		visit(root.Left)
		inorder = append(inorder, root.Val)
		visit(root.Right)
	}
	visit(root)
	if len(inorder) < 2 {
		return true
	}
	for i := 1; i < len(inorder); i++ {
		if inorder[i-1] >= inorder[i] {
			return false
		}
	}
	return true
}
