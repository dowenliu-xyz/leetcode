package lc0007_2

import "math"

func reverse(x int) int {
	maxQ, maxR, minQ, minR := math.MaxInt32/10, math.MaxInt32%10, math.MinInt32/10, math.MinInt32%10
	ans := 0
	for x != 0 {
		r := x % 10
		if x < 0 && (ans < minQ || (ans == minQ && r < minR)) {
			return 0
		}
		if x > 0 && (ans > maxQ || (ans == maxQ && r > maxR)) {
			return 0
		}
		ans = ans*10 + r
		x /= 10
	}
	return ans
}
