package main

func longestPalindrome(s string) string {
	//return solution1(s)
	//return solution2(s)
	return solution3(s)
}

func solution1(s string) string {
	// 暴力
	runes := []rune(s)
	maxLen, res := 1, runes[:1]
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			r := runes[i : j+1]
			if isPalindrome(r) && j-i+1 > maxLen {
				maxLen = j - i + 1
				res = r
			}
		}
	}
	return string(res)
}

func isPalindrome(runes []rune) bool {
	l := len(runes)
	for i := 0; i < l/2; i++ {
		if runes[i] != runes[l-1-i] {
			return false
		}
	}
	return true
}

func solution2(s string) string {
	// dp
	runes := []rune(s)
	n := len(runes)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	maxLen, begin := 1, 0
	for L := 2; L <= n; L++ { // 枚举长度
		for i := 0; i < n; i++ { // 枚举左边界
			j := L + i - 1 // 对应右边界
			if j >= n {    // 右边界越界
				break
			}
			if runes[i] == runes[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return string(runes[begin : begin+maxLen])
}

func solution3(s string) string {
	// 中心扩展法
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
