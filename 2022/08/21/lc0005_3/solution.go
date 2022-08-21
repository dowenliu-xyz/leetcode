package main

func longestPalindrome(s string) string {
	//return solution1(s)
	return solution2(s)
	return solution3(s)
}

func solution1(s string) string {
	// 暴力枚举
	if len(s) < 2 {
		return s
	}
	begin, maxLen := 0, 1
	for i := 0; i < len(s); i++ {
		for j := i + maxLen; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				begin, maxLen = i, j-i+1
			}
		}
	}
	return s[begin : begin+maxLen]
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func solution2(s string) string {
	// dp
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	begin, maxLen := 0, 1
	// 枚举子串长度
	for L := 2; L <= len(s); L++ {
		// 枚举子串窗口左起点
		for left := 0; left < len(s); left++ {
			// 确定子串窗口右截止点
			// right - left + 1 = L
			right := left + L - 1
			if right >= len(s) {
				// 子串窗口右侧越界
				break
			}
			if s[left] != s[right] {
				// 子串窗口右移
				continue
			}
			if L < 3 {
				dp[left][right] = true
			} else {
				dp[left][right] = dp[left+1][right-1]
			}
			if dp[left][right] && L > maxLen {
				begin, maxLen = left, L
			}
		}
	}
	return s[begin : begin+maxLen]
}

func solution3(s string) string {
	// 中心扩展
	if len(s) < 2 {
		return s
	}
	begin, maxLen := 0, 1
	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r1-l1+1 > maxLen {
			begin, maxLen = l1, r1-l1+1
		}
		if r2-l2+1 > maxLen {
			begin, maxLen = l2, r2-l2+1
		}
	}
	return s[begin : begin+maxLen]
}

func expandAroundCenter(s string, l, r int) (int, int) {
	for ; l >= 0 && l < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}
