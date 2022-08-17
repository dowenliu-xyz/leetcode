package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//return solution1(root, p, q)
	return solution2(root, p, q)
}

func solution1(root, p, q *TreeNode) *TreeNode {
	// dfs
	if root == nil {
		return nil
	}
	// 1) p, q 分别位于 root 左右子树中
	// 2) p, q 位于 root 的一侧子树中
	// 3) p, q 其中之一就是 root
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
	// bfs 记录父节点
	if root == nil {
		return nil
	}
	pmp := map[int]*TreeNode{}
	dfs(root, pmp)
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

func dfs(root *TreeNode, pmp map[int]*TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			pmp[node.Left.Val] = node
			q = append(q, node.Left)
		}
		if node.Right != nil {
			pmp[node.Right.Val] = node
			q = append(q, node.Right)
		}
	}
}
