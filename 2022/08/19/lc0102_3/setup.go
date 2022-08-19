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

func formatResult(result [][]int) string {
	lines := make([]string, 0, len(result))
	for _, vals := range result {
		line := make([]string, 0, len(vals))
		for _, val := range vals {
			line = append(line, strconv.Itoa(val))
		}
		lines = append(lines, fmt.Sprintf("[%s]", strings.Join(line, ",")))
	}
	return fmt.Sprintf("[%s]", strings.Join(lines, ","))
}
