package lc0069_1

func mySqrt(x int) int {
	return binarySearch(0, x, x)
}

func binarySearch(l, h, t int) int {
	m := l + (h-l)>>1
	ms := m * m
	if ms == t {
		return m
	} else if ms > t {
		return binarySearch(l, m-1, t)
	}
	ns := (m + 1) * (m + 1)
	if ns > t {
		return m
	}
	return binarySearch(m+1, h, t)
}
