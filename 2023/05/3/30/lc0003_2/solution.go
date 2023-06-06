package lc0003_2

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	ans, l, mem := 1, 0, make(map[byte]bool)
	for r := 0; r < n; r++ {
		for mem[s[r]] && l < r {
			delete(mem, s[l])
			l++
		}
		mem[s[r]] = true
		if ans < r-l+1 {
			ans = r - l + 1
		}
	}
	return ans
}
