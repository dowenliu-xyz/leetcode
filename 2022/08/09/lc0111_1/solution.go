package lc0111_1

import "math"

func minDepth(root *TreeNode) int {
	//return solution1(root)
	//return solution2(root)
	return solution3(root)
}

func solution1(root *TreeNode) int {
	// 递归，深度优先搜索
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	childMin := math.MaxInt
	if root.Left != nil {
		childMin = min(childMin, minDepth(root.Left))
	}
	if root.Right != nil {
		childMin = min(childMin, minDepth(root.Right))
	}
	return childMin + 1
}

func solution2(root *TreeNode) int {
	// 简化递归，深度优先搜索
	if root == nil {
		return 0
	}
	ml, mr := minDepth(root.Left), minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return ml + mr + 1
	}
	return min(ml, mr) + 1
}

func solution3(root *TreeNode) int {
	// 迭代，广度优先搜索
	if root == nil {
		return 0
	}
	m := 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			if node.Left == nil && node.Right == nil {
				return m
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		m++
		q = p
	}
	return m
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
