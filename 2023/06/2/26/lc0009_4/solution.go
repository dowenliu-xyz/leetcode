package lc0009

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x != 0 && x%10 == 0 {
		return false
	}
	maxQ, maxR := math.MaxInt/10, math.MaxInt%10
	cur := 0
	for x > cur {
		tmp := x % 10
		if cur > maxQ || (cur == maxQ && tmp > maxR) {
			return false
		}
		cur = cur*10 + tmp
		x /= 10
	}
	return cur == x || cur/10 == x
}
