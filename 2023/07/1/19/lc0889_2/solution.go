package lc0889

import . "github.com/dowenliu-xyz/leetcode/internal/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if n == 1 {
		return root
	}
	leftN := 0
	for postorder[leftN] != preorder[1] {
		leftN++
	}
	leftN++
	root.Left = constructFromPrePost(preorder[1:1+leftN], postorder[:leftN])
	root.Right = constructFromPrePost(preorder[1+leftN:], postorder[leftN:n-1])
	return root
}
