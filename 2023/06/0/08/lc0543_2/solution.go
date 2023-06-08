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
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := height(root.Left), height(root.Right)
		ans = max(ans, l+r)
		return max(l, r) + 1
	}
	height(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
