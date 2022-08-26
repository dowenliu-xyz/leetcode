package main

import "strings"

func isValid(s string) bool {
	//return solution1(s)
	return solution2(s)
}

func solution1(s string) bool {
	// 循环替换
	for len(s) > 0 {
		l := len(s)
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
		if len(s) == l {
			return false
		}
	}
	return true
}

func solution2(s string) bool {
	// 使用栈
	var stack []byte
	pair := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < len(s); i++ {
		if c, ok := pair[s[i]]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || s[i] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
