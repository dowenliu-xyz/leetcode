package main

func isCompleteTree(root *TreeNode) bool {
	// 层序遍历，遇到 nil 后不能再遇到任何节点
	q := []*TreeNode{root}
	seeNil := false
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node != nil && seeNil {
			return false
		}
		if node == nil {
			seeNil = true
			continue
		}
		q = append(q, node.Left, node.Right)
	}
	return true
}
