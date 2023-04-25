package lc0101_5

func isSymmetric(root *TreeNode) bool {
	//return solution1(root, root)
	return solution2(root, root)
}

func solution1(p, q *TreeNode) bool {
	// dfs
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && solution1(p.Left, q.Right) && solution1(p.Right, q.Left)
}

func solution2(p, q *TreeNode) bool {
	// bfs
	pq, qq := []*TreeNode{p}, []*TreeNode{q}
	for len(pq) > 0 && len(qq) > 0 {
		p, q = pq[0], qq[0]
		pq, qq = pq[1:], qq[1:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		pq, qq = append(pq, p.Left, p.Right), append(qq, q.Right, q.Left)
	}
	return len(pq) == 0 && len(qq) == 0
}
