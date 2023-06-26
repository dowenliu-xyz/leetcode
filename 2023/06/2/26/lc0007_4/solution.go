package lc0007

import "math"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	ans := 0
	if x > 0 {
		maxQ, maxR := math.MaxInt32/10, math.MaxInt32%10
		for x != 0 {
			tmp := x % 10
			if ans > maxQ || (ans == maxQ && tmp > maxR) {
				return 0
			}
			ans = ans*10 + tmp
			x /= 10
		}
	} else {
		minQ, minR := math.MinInt32/10, math.MinInt32%10
		for x != 0 {
			tmp := x % 10
			if ans < minQ || (ans == minQ && tmp < minR) {
				return 0
			}
			ans = ans*10 + tmp
			x /= 10
		}
	}
	return ans
}
