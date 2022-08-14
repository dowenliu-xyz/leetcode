package main

func longestPalindrome(s string) string {
	//return solution1(s)
	//return solution2(s)
	//return solution3(s)
	return solution4(s)
}

func solution1(s string) string {
	// 暴力枚举
	if len(s) < 2 {
		return s
	}
	maxLen, begin := 1, 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func solution2(s string) string {
	// 还是暴力枚举，但稍微优化
	if len(s) < 2 {
		return s
	}
	maxLen, begin := 1, 0
	for i := 0; i < len(s); i++ {
		// 我们只检查长度比已经最大长度长的子串即可
		for j := i + maxLen; j < len(s); j++ {
			// j-i+1 必定大于 maxLen
			if isPalindrome(s[i : j+1]) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func solution3(s string) string {
	// dp。对于这个问题，dp 能解决问题，但dp数组的利用率不高，所以效率提升有限
	sLen := len(s)
	if sLen < 2 {
		return s
	}
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
		dp[i][i] = true
	}
	maxLen, begin := 1, 0
	// 枚举可能长度。注意这里可能长度可能等于字符串长度，不要写成严格小于
	for L := 2; L <= sLen; L++ {
		// 枚举左边界
		for left := 0; left < sLen; left++ {
			// 确定左边界和子串长度，可计算得到右边界
			// right - left + 1 = L
			right := L + left - 1
			if right >= sLen {
				// 右边界越界，无效
				break
			}
			if s[left] != s[right] {
				// 不是回文串
				continue // 如果是 break 会导致可能以后面位置作为左边界的回文串被跳过
			}
			if L < 3 {
				// 1) a
				// 2) aa
				// 3) aba
				dp[left][right] = true
			} else {
				dp[left][right] = dp[left+1][right-1]
			}
			if dp[left][right] && L > maxLen {
				maxLen = L
				begin = left
			}
		}
	}
	return s[begin : begin+maxLen]
}

func solution4(s string) string {
	// 中心扩展法
	sLen := len(s)
	if sLen < 2 {
		return s
	}
	maxLen, begin := 1, 0
	for i := 0; i < sLen; i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1+1 > maxLen {
			maxLen = right1 - left1 + 1
			begin = left1
		}
		if right2-left2+1 > maxLen {
			maxLen = right2 - left2 + 1
			begin = left2
		}
	}
	return s[begin : begin+maxLen]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
