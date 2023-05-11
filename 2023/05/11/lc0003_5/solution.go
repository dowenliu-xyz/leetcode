package lc0003_5

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}
	ans, t, h, mem := 1, 0, 0, map[byte]bool{}
	for h < l {
		for mem[s[h]] && t < h {
			delete(mem, s[t])
			t++
		}
		mem[s[h]] = true
		h++
		if ans < h-t {
			ans = h - t
		}
	}
	return ans
}
