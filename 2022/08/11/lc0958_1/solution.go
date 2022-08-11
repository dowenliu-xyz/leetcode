package lc0958_1

func isCompleteTree(root *TreeNode) bool {
	//return solution1(root)
	return solution2(root)
}

func solution1(root *TreeNode) bool {
	// 层序遍历，广度优先搜索
	// 对每一层，每一个节点：
	// 1）如果前面出现过 nil 节点，那么这个节点必需也是 nil
	// 2）如果前面的非 nil 节点有子节点，那么这个节点不能是 nil
	if root == nil {
		return true
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		seenChild, seenNil := false, false
		var p []*TreeNode
		for _, node := range q {
			if node.Left == nil {
				if seenChild {
					return false
				}
				seenNil = true
			} else {
				if seenNil {
					return false
				}
				p = append(p, node.Left)
				if node.Left.Left != nil || node.Left.Right != nil {
					seenChild = true
				}
			}
			if node.Right == nil {
				if seenChild {
					return false
				}
				seenNil = true
			} else {
				if seenNil {
					return false
				}
				p = append(p, node.Right)
				if node.Right.Left != nil || node.Right.Right != nil {
					seenChild = true
				}
			}
		}
		q = p
	}
	return true
}

func solution2(root *TreeNode) bool {
	// 层序遍历，广度优先搜索
	// 按照定义，可以知道，层序遍历要么遇不到 nil，要么遇到的nil是整个层序遍历的结束
	// 如果遍历过程中遇到 nil ，之后再遇到任何节点都说明不是完全二叉树
	if root == nil {
		return true
	}
	q, seenNil := []*TreeNode{root}, false
	for len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			if node.Left == nil {
				seenNil = true
			} else {
				if seenNil {
					return false
				}
				p = append(p, node.Left)
			}
			if node.Right == nil {
				seenNil = true
			} else {
				if seenNil {
					return false
				}
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return true
}
