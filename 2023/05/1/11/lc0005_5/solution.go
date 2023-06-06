package lc0005_5

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	max, ex, left, right := 1, 0, 0, 0
	for i := 0; i < l-1; i++ {
		exLen := tryAsCenter(s, i, i)
		subLen := exLen<<1 + 1
		if max < subLen {
			max, ex, left, right = subLen, exLen, i, i
		}
		exLen = tryAsCenter(s, i, i+1)
		subLen = exLen<<1 + 2
		if max < subLen {
			max, ex, left, right = subLen, exLen, i, i+1
		}
	}
	return s[left-ex : right+ex+1]
}

func tryAsCenter(s string, l, r int) int {
	ans := -1
	for l >= 0 && r < len(s) && s[l] == s[r] {
		ans++
		l--
		r++
	}
	return ans
}
