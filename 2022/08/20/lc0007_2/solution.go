package main

import "math"

func reverse(x int) int {
	max, maxR := math.MaxInt32/10, math.MaxInt32%10
	min, minR := math.MinInt32/10, math.MinInt32%10
	reverted := 0
	for x != 0 {
		tmp := x % 10
		if reverted > max || (reverted == max && tmp > maxR) {
			return 0
		}
		if reverted < min || (reverted == min && tmp < minR) {
			return 0
		}
		reverted = reverted*10 + tmp
		x /= 10
	}
	return reverted
}
