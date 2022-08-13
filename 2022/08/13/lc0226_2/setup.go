package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func parseStrings(input string) []string {
	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	parts := strings.Split(input, ",")
	res := make([]string, 0, len(parts))
	for _, s := range parts {
		trim := strings.TrimSpace(s)
		if "" != trim {
			res = append(res, trim)
		}
	}
	return res
}

func levelOrderRestore(input []string) *TreeNode {
	if len(input) == 0 {
		return nil
	}
	if strings.EqualFold("null", input[0]) {
		return nil
	}
	val, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}
	root := &TreeNode{Val: val}
	q := []*TreeNode{root}
	input = input[1:]
	for len(input) > 0 && len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			if len(input) == 0 {
				break
			}
			if strings.EqualFold("null", input[0]) {
				node.Left = nil
			} else {
				val, err := strconv.Atoi(input[0])
				if err != nil {
					panic(err)
				}
				node.Left = &TreeNode{Val: val}
				p = append(p, node.Left)
			}
			if len(input) == 1 {
				input = input[1:]
				break
			}
			if strings.EqualFold("null", input[1]) {
				node.Right = nil
			} else {
				val, err := strconv.Atoi(input[1])
				if err != nil {
					panic(err)
				}
				node.Right = &TreeNode{Val: val}
				p = append(p, node.Right)
			}
			input = input[2:]
		}
		q = p
	}
	return root
}

func levelOrderFormat(root *TreeNode) string {
	var nodes []string
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			nodes = append(nodes, "null")
			continue
		}
		nodes = append(nodes, strconv.Itoa(node.Val))
		q = append(q, node.Left, node.Right)
	}
	result := strings.Join(nodes, ",")
	for {
		result = strings.TrimRight(result, ",")
		if !strings.HasSuffix(result, "null") {
			break
		}
		result = strings.TrimRight(result, "null")
	}
	return fmt.Sprintf("[%s]", result)
}
