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
	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		node1, node2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		queue1 = append(queue1, node1.Left, node1.Right)
		queue2 = append(queue2, node2.Left, node2.Right)
	}
	return len(queue1) == 0 && len(queue2) == 0
}
