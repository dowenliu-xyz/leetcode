package lc0226_5

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
	ans := root
	q := []*TreeNode{root}
	for len(q) > 0 {
		root = q[0]
		q = q[1:]
		if root == nil {
			continue
		}
		root.Left, root.Right = root.Right, root.Left
		q = append(q, root.Left, root.Right)
	}
	return ans
}
