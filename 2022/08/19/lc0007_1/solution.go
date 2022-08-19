package main

import "math"

func reverse(x int) int {
	maxAllow, minAllow := math.MaxInt32/10, math.MinInt32/10
	maxR, minR := math.MaxInt32%10, math.MinInt32%10
	var ans int
	for x != 0 {
		tmp := x % 10
		if ans > maxAllow || (ans == maxAllow && tmp > maxR) {
			return 0
		}
		if ans < minAllow || (ans == minAllow && tmp < minR) {
			return 0
		}
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}
