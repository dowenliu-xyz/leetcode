package lc0072

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		// 如何任一单词为空白，编辑距离就是另一单词长度
		return m + n
	}
	// dp 数组: dp[i][j] 表示 word1[:i] 到 word2[:j] 的编辑距离
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // word1[:i] 到空白的编辑距离为 word1[:i] 的长度，即 i
	}
	for i := range dp[0] {
		dp[0][i] = i // 空白到 word2[:i] 的编辑距离为 word2[:i] 的长度，即 i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				// 编辑距离不变
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 需要至少一次替换、新增或删除
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
