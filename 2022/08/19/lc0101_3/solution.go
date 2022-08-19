package main

func isSymmetric(root *TreeNode) bool {
	//return solution1(root, root)
	return solution2(root, root)
}

func solution1(root1, root2 *TreeNode) bool {
	// dfs
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return solution1(root1.Left, root2.Right) && solution1(root1.Right, root2.Left)
}

func solution2(root1, root2 *TreeNode) bool {
	// bfs
	q := [][2]*TreeNode{{root1, root2}}
	for len(q) > 0 {
		node1, node2 := q[0][0], q[0][1]
		q = q[1:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		q = append(q, [2]*TreeNode{node1.Left, node2.Right}, [2]*TreeNode{node1.Right, node2.Left})
	}
	return true
}
