package lc0009_3

import "math"

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	maxQ, maxR := math.MaxInt32/10, math.MaxInt32%10
	ans := 0
	for ans < x {
		tmp := x % 10
		if ans > maxQ || (ans == maxQ && tmp > maxR) {
			return false
		}
		ans = ans*10 + tmp
		x /= 10
	}
	return ans == x || ans/10 == x
}
