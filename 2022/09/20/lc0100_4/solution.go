package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	//return solution1(p, q)
	return solution2(p, q)
}

func solution1(p, q *TreeNode) bool {
	// dfs
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return solution1(p.Left, q.Left) && solution1(p.Right, q.Right)
}

func solution2(p, q *TreeNode) bool {
	// bfs
	q1, q2 := []*TreeNode{p}, []*TreeNode{q}
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
		q2 = append(q2, n2.Left, n2.Right)
	}
	return len(q1) == 0 && len(q2) == 0
}
