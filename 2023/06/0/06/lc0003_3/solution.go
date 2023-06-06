package lc0003

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	l, r, mem := 0, 0, map[byte]bool{}
	ans := 1
	for r < n {
		for mem[s[r]] {
			delete(mem, s[l])
			l++
		}
		mem[s[r]] = true
		r++
		ans = max(ans, r-l)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
