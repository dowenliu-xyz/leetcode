package lc0322

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1) // dp[i] 表示找零 i 需要的最少硬币数
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i < coin {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
