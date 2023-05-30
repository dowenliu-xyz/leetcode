package lc0102_2

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
		var line []int
		for _, node := range q {
			line = append(line, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		ans = append(ans, line)
		q = p
	}
	return ans
}
