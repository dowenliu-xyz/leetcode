package main

import (
	"math"
	"strconv"
)

func isPalindrome(x int) bool {
	//return solution1(x)
	return solution2(x)
}

func solution1(x int) bool {
	// 转字符串，再判断是否回文串
	s := strconv.Itoa(x)
	for i := 0; i < len(s)>>1; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func solution2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	m, n := math.MaxInt32/10, math.MaxInt32%10
	cur := 0
	for cur < x {
		tmp := x % 10
		if cur > m || (cur == m && tmp > n) { // 会溢出
			return false
		}
		cur = cur*10 + tmp
		x /= 10
	}
	return cur == x || cur/10 == x
}
