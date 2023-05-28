package lc0474_3

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	// dp 数组。每考察 strs 的一个元素后，dp[i][j] 表示最多允许 i 个 0 、 j 个 1 时可包含的最大子元素数量（很有可能用不完这些 0 、1 限制）
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for i := m; i >= 0; i-- { // 要包括 i 等于 0 的情况，比如允许 0 个 0 、3 个 1 时
			for j := n; j >= 0; j-- { // 同样要包括 j 等于 0 的情况
				if i < zeros || j < ones {
					continue
				}
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
