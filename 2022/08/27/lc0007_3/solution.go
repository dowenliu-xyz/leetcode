package main

import "math"

func reverse(x int) int {
	maxCur, maxTmp := math.MaxInt32/10, math.MaxInt32%10
	minCur, minTmp := math.MinInt32/10, math.MinInt32%10
	cur := 0
	for x != 0 {
		tmp := x % 10
		if x < 0 && (cur < minCur || (cur == minCur && tmp < minTmp)) {
			return 0
		}
		if x > 0 && (cur > maxCur || (cur == maxCur && tmp > maxTmp)) {
			return 0
		}
		cur = cur*10 + tmp
		x /= 10
	}
	return cur
}
