package lc0010_2

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1) // dp 数组，dp[i][j] 表示 s[:i] 能否被 p[:j] 匹配
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	// 处理 dp 首行
	dp[0][0] = true // dp[0][0] 表示空白 s 能被空白 p 匹配
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2] // p[j-2:j] 可能匹配空白，所以 p[:j] 能否匹配空白取决于 p[:j-2]
		}
	}
	// 依次考察 s[:i] 能否被 p[:j] 匹配
	for i := 1; i <= m; i++ { // s 的第 i 个字符 s[i-1]，前 i 个字符组成的子串 s[:i]
		for j := 1; j <= n; j++ { // p 的第 j 个字符 p[j-1]，前 j 个字符组成的子串 p[:j]
			if p[j-1] == '*' {
				// p[j-1] 为 * ，则在确定 j >= 2 且 p[j-2] 为字母或'.'
				// 考察 p[j-2:j] 匹配到空白的情况，即检查 s[:i] 是否能被 p[:j-2] 匹配
				dp[i][j] = dp[i][j] || dp[i][j-2]
				// 考察 p[:j] 是否已经能匹配 s[:i-1]，如果已经能匹配，p[j-1] 为 * 表示 p[:j] 能继续匹配 s[:i]
				if match(s[i-1], p[j-2]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if match(s[i-1], p[j-1]) {
				// p[j-1] 不是 * 且能匹配 s[i-1]，则 s[:i] 能否被 p[:j] 匹配就看 s[:i-1] 能否被 p[:j-1] 匹配
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
