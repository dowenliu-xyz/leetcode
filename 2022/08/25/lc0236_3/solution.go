package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//return solution1(root, p, q)
	return solution2(root, p, q)
}

func solution1(root, p, q *TreeNode) *TreeNode {
	// dfs
	// p, q 只能是以下三种情形之一
	// 1） p 或 q 是 root
	// 2） p 或 q 是对方的先祖
	// 3） p 和 q 分别处于 root 的左右子树
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left, right := solution1(root.Left, p, q), solution1(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func solution2(root, p, q *TreeNode) *TreeNode {
	// 记录父节点
	if root == nil {
		return nil
	}
	pmp := map[int]*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			pmp[node.Left.Val] = node
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			pmp[node.Right.Val] = node
			queue = append(queue, node.Right)
		}
	}
	visited := map[int]bool{}
	for p != nil {
		visited[p.Val] = true
		p = pmp[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = pmp[q.Val]
	}
	return root
}
