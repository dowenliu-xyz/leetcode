package lc0010_1

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1) // s[0:i] 是否与 p[0:j] 匹配
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	// 处理首行，是否可以匹配空白
	dp[0][0] = true // 空白 p 一定匹配空白 s
	for i := 2; i <= n; i++ {
		// 处理完全由 * 模式的情况：a*、.*、a*b* 等
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}
	for i := 1; i <= m; i++ { // s 第 i 个字符 s[i-1]
		for j := 1; j <= n; j++ { // p 第 j 个字符 p[j-1]
			if p[j-1] == '*' { // p[j-1] 是 *，则保证 j >= 2 且 p[j-2] 是字母或 .
				dp[i][j] = dp[i][j] || dp[i][j-2] // 可能 x* 匹配到 0 个字符，x 为 字母或 .
				if match(s[i-1], p[j-2]) {        // s[i-1] 与 p[j-2] 匹配
					dp[i][j] = dp[i][j] || dp[i-1][j] // 如果 s[:i-1] 与 p[:j] 匹配，s[i-1] 与 p[j-2] 匹配，那么 s[:i] 与 p[:j] 继续匹配
				}
			} else if match(s[i-1], p[j-1]) { // s[i-1] 与 p[j-1] 匹配
				dp[i][j] = dp[i][j] || dp[i-1][j-1] // 那么 s[:i] 与 p[:j] 是否匹配就取决于 s[:i-1] 与 s[:j-1] 是否匹配
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
