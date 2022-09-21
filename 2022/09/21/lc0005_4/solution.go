package main

func longestPalindrome(s string) string {
	//return solution1(s)
	//return solution2(s)
	return solution3(s)
}

func solution1(s string) string {
	// 暴力
	n := len(s)
	if n < 2 {
		return s
	}
	begin, maxLen := 0, 1
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if j-i+1 > maxLen && isPalindrome(s, i, j) {
				begin, maxLen = i, j-i+1
			}
		}
	}
	return s[begin : begin+maxLen]
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func solution2(s string) string {
	// dp
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	begin, maxLen := 0, 1
	for L := 2; L <= n; L++ {
		for i := 0; i < n; i++ {
			j := i + L - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				continue
			}
			if L < 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && L > maxLen {
				begin, maxLen = i, L
			}
		}
	}
	return s[begin : begin+maxLen]
}

func solution3(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	begin, maxLen := 0, 1
	for i := 0; i < n; i++ {
		l, r := expandAroundCenter(s, i, i)
		if r-l+1 > maxLen {
			begin, maxLen = l, r-l+1
		}
		l, r = expandAroundCenter(s, i, i+1)
		if r-l+1 > maxLen {
			begin, maxLen = l, r-l+1
		}
	}
	return s[begin : begin+maxLen]
}

func expandAroundCenter(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}
