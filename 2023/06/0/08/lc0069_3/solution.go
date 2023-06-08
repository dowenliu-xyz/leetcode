package lc0069

func mySqrt(x int) int {
	if x < 0 {
		return -1
	}
	l, r := 0, x
	for l <= r {
		m := l + (r-l)>>1
		s := m * m
		if x < s {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r
}
