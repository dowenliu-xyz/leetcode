package main

func isSymmetric(root *TreeNode) bool {
	//return solution1(root, root)
	return solution2(root, root)
}

func solution1(left, right *TreeNode) bool {
	// dfs
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return solution1(left.Left, right.Right) && solution1(left.Right, right.Left)
}

func solution2(left, right *TreeNode) bool {
	// bfs
	q1, q2 := []*TreeNode{left}, []*TreeNode{right}
	for len(q1) > 0 && len(q2) > 0 {
		n1, n2 := q1[0], q2[0]
		q1, q2 = q1[1:], q2[1:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		q1 = append(q1, n1.Left, n1.Right)
		q2 = append(q2, n2.Right, n2.Left)
	}
	return len(q1) == 0 && len(q2) == 0
}
