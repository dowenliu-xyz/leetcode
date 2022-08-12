package lc0002_2

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func readInput(input string) *ListNode {
	ints := readInts(input)
	return intList(ints)
}

func readInts(input string) []int {
	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	parts := strings.Split(input, ",")
	res := make([]int, 0, len(parts))
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}

func intList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	node := head
	for i := 1; i < len(vals); i++ {
		node.Next = &ListNode{Val: vals[i]}
		node = node.Next
	}
	return head
}

func formatList(l *ListNode) string {
	builder := &strings.Builder{}
	builder.WriteRune('[')
	for l != nil {
		builder.WriteString(strconv.Itoa(l.Val))
		l = l.Next
		if l != nil {
			builder.WriteRune(',')
		}
	}
	builder.WriteRune(']')
	return builder.String()
}
