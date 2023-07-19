package lc0094

import . "github.com/dowenliu-xyz/leetcode/internal/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) []int {
	var ans []int
	var visit func(root *TreeNode)
	visit = func(root *TreeNode) {
		if root == nil {
			return
		}
		visit(root.Left)
		ans = append(ans, root.Val)
		visit(root.Right)
	}
	visit(root)
	return ans
}

func solution2(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}
