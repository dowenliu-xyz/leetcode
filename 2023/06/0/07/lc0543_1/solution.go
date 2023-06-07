package lc0543

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := height(node.Left), height(node.Right)
		ans = max(ans, l+r)
		return max(l, r) + 1
	}
	height(root)
	return ans
}
