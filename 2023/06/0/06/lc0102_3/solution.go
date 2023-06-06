package lc0102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		var level []int
		for _, node := range q {
			level = append(level, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		ans = append(ans, level)
		q = p
	}
	return ans
}
