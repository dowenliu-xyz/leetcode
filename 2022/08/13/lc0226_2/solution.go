package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	//return solution1(root)
	//return solution2(root)
	//return solution3(root)
	return solution4(root)
}

func solution1(root *TreeNode) *TreeNode {
	// 原地操作，递归 dfs
	if root == nil {
		return root
	}
	root.Left, root.Right = solution1(root.Right), solution1(root.Left)
	return root
}

func solution2(root *TreeNode) *TreeNode {
	// 复制，递归 dfs
	if root == nil {
		return nil
	}
	res := &TreeNode{Val: root.Val}
	res.Left, res.Right = solution2(root.Right), solution2(root.Left)
	return res
}

func solution3(root *TreeNode) *TreeNode {
	// 原地操作，迭代 bfs
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return root
}

func solution4(root *TreeNode) *TreeNode {
	// 复制。迭代 bfs
	if root == nil {
		return nil
	}
	res := &TreeNode{
		Val:   root.Val,
		Left:  root.Left,
		Right: root.Right,
	}
	q := []*TreeNode{res}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		var left, right *TreeNode
		if node.Right != nil {
			left = &TreeNode{
				Val:   node.Right.Val,
				Left:  node.Right.Left,
				Right: node.Right.Right,
			}
			q = append(q, left)
		}
		if node.Left != nil {
			right = &TreeNode{
				Val:   node.Left.Val,
				Left:  node.Left.Left,
				Right: node.Left.Right,
			}
			q = append(q, right)
		}
		node.Left, node.Right = left, right
	}
	return res
}
