package main

func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	seenNil := false
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			seenNil = true
			continue
		}
		if seenNil {
			return false
		}
		q = append(q, node.Left, node.Right)
	}
	return true
}
