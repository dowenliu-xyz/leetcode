package lc0474_1

import "strings"

func findMaxForm(strs []string, m int, n int) int {
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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
