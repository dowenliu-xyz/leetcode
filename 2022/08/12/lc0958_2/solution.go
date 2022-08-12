package lc0958_2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	// 根据定义，可按bfs 层序遍历。当遇到 nil 后，如果再遇到非 nil 值，则说明不是完全二叉树
	q, seenNil := []*TreeNode{root}, false

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			seenNil = true
		} else {
			if seenNil {
				return false
			}
			q = append(q, node.Left, node.Right)
		}
	}
	return true
}
