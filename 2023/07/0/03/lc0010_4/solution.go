package lc0010

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// dp[i][j] 表示 s[:i] 能否被 p[:j] 匹配
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(s[i-1], p[j-2]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if match(s[i-1], p[j-1]) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func match(a, b byte) bool {
	if a == b {
		return true
	}
	if b == '.' {
		return true
	}
	return false
}
