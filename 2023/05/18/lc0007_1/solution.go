package lc0007_1

import "math"

func reverse(x int) int {
	maxQ, maxR, minQ, minR := math.MaxInt32/10, math.MaxInt32%10, math.MinInt32/10, math.MinInt32%10
	ans := 0
	for x != 0 {
		tmp := x % 10
		if x < 0 && (ans < minQ || (ans == minQ && tmp < minR)) {
			return 0
		}
		if x > 0 && (ans > maxQ || (ans == maxQ && tmp > maxR)) {
			return 0
		}
		ans = ans*10 + tmp
		x /= 10
	}
	return ans
}
