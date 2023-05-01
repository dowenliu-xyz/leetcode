package lc0120_1

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp1 := make([]int, n)
	dp1[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp1[i] = dp1[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp1[j] = min(dp1[j], dp1[j-1]) + triangle[i][j]
		}
		dp1[0] = dp1[0] + triangle[i][0]
	}
	ans := math.MaxInt
	for i := 0; i < n; i++ {
		if dp1[i] < ans {
			ans = dp1[i]
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
