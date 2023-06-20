package lc0096

func numTrees(n int) int {
	// 特殊值，基本情景。0个或1个节点的情况都只有一种
	if n <= 1 {
		return 1
	}
	// dp 数组。dp[i] 表示由i个节点组成的种数
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1 // 初始化特殊值
	for i := 2; i <= n; i++ {
		// dp[i] = dp[0]*dp[i-1] + dp[1]*dp[i-2] + ... + dp[j-1]*dp[i-j], j: 1 -> i
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
