package lc0440

func findKthNumber(n int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	getCount := func(prefix int) int {
		cur, next, count := prefix, prefix+1, 0
		for cur <= n {
			count += min(n+1, next) - cur
			cur *= 10
			next *= 10
		}
		return count
	}
	p, prefix := 1, 1
	for p < k {
		count := getCount(prefix)
		if p+count > k {
			prefix *= 10
			p++
		} else {
			prefix++
			p += count
		}
	}
	return prefix
}
