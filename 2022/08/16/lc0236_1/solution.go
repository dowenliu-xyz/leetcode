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
	// 递归
	// p、q 要么在分别 root 的左右子树中，要么在一侧子树中，要么 root 就是 p、q 之一。
	// 如果 p、q 不是分别在 root 的左右子树中，那么必然 p、q 中的一个是另一个节点的先祖
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	fromLeft := solution1(root.Left, p, q)
	fromRight := solution1(root.Right, p, q)
	if fromLeft != nil && fromRight != nil {
		return root
	}
	if fromLeft == nil {
		return fromRight
	}
	return fromLeft
}

func solution2(root, p, q *TreeNode) *TreeNode {
	// 记录父节点
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
	return nil
}

func dfs(root *TreeNode, pmp map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		pmp[root.Left.Val] = root
		dfs(root.Left, pmp)
	}
	if root.Right != nil {
		pmp[root.Right.Val] = root
		dfs(root.Right, pmp)
	}
}
