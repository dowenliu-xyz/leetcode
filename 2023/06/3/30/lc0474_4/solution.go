package lc0474

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j] 表示在最多有 i 个 0 和 j 个 1 时，最大子集长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 遍历元素，每个元素只能被使用一次
	for _, str := range strs {
		zeros := strings.Count(str, "0")
		ones := len(str) - zeros
		for i := m; i >= 0; i-- {
			for j := n; j >= 0; j-- {
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
