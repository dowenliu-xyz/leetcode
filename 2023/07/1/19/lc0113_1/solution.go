package lc0113

import . "github.com/dowenliu-xyz/leetcode/internal/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var dfs func(path []int, sum int, root *TreeNode)
	dfs = func(path []int, sum int, root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		sum += root.Val
		if root.Left == nil && root.Right == nil && sum == targetSum {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		if root.Left != nil {
			dfs(path, sum, root.Left)
		}
		if root.Right != nil {
			dfs(path, sum, root.Right)
		}
	}
	dfs(nil, 0, root)
	return ans
}
