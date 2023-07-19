package lc0701

import . "github.com/dowenliu-xyz/leetcode/internal/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	//return solution1(root, val)
	return solution2(root, val)
}

func solution1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = solution1(root.Left, val)
	} else {
		root.Right = solution1(root.Right, val)
	}
	return root
}

func solution2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	p := root
	for {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				return root
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				return root
			}
			p = p.Right
		}
	}
}
