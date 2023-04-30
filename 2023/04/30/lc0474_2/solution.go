package lc0474_2

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	// dp 数组。滚动数组，每考察一个元素后, dp[j][k] 表示达到 j 个 0 、 k 个 1 的最大子集合数。如果达不到， 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for j := m; j >= 0; j-- {
			for k := n; k >= 0; k-- {
				if j >= zeros && k >= ones {
					dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
				}
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
