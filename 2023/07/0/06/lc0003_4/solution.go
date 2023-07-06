package lc0003

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	seen := make([]bool, 256)
	ans := 1
	l, r := 0, 0
	for ; r < n; r++ {
		for seen[s[r]] {
			seen[s[l]] = false
			l++
		}
		seen[s[r]] = true
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
