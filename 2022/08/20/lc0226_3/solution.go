package main

func invertTree(root *TreeNode) *TreeNode {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) *TreeNode {
	// dfs
	if root == nil {
		return nil
	}
	root.Left, root.Right = solution1(root.Right), solution1(root.Left)
	return root
}

func solution2(root *TreeNode) *TreeNode {
	// bfs
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		node.Left, node.Right = node.Right, node.Left
		q = append(q, node.Left, node.Right)
	}
	return root
}
