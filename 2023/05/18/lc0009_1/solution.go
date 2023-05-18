package lc0009_1

import "math"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	m, n := math.MaxInt32/10, math.MaxInt32%10
	cur := 0
	for cur < x {
		tmp := x % 10
		if cur > m || (cur == m && tmp > n) {
			return false
		}
		cur = cur*10 + tmp
		x /= 10
	}
	return cur == x || cur/10 == x
}
