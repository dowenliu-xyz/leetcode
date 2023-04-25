package lc0100_5

func isSameTree(p *TreeNode, q *TreeNode) bool {
	//return solution1(p, q)
	return solution2(p, q)
}

func solution1(p *TreeNode, q *TreeNode) bool {
	// dfs
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && solution1(p.Left, q.Left) && solution1(p.Right, q.Right)
}

func solution2(p, q *TreeNode) bool {
	// bfs
	pq, qq := []*TreeNode{p}, []*TreeNode{q}
	for len(pq) != 0 && len(qq) != 0 {
		p, q := pq[0], qq[0]
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
		pq, qq = append(pq, p.Left, p.Right), append(qq, q.Left, q.Right)
	}
	return len(pq) == 0 && len(qq) == 0
}
