package lc0440

func findKthNumber(n int, k int) int {
	getCount := func(prefix int) int {
		next, count := prefix+1, 0
		for prefix <= n {
			count += min(n+1, next) - prefix
			prefix, next = prefix*10, next*10
		}
		return count
	}
	p, prefix := 1, 1
	for p < k {
		count := getCount(prefix)
		if p+count <= k {
			prefix++
			p += count
		} else {
			prefix *= 10
			p++
		}
	}
	return prefix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
