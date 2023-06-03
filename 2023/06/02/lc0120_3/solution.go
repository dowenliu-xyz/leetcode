package lc0120_3

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = triangle[i][j] + min(dp[j], dp[j-1])
		}
		dp[0] += triangle[i][0]
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		ans = min(ans, dp[i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
