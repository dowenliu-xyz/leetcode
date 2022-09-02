package main

func maxProfit(prices []int) int {
	// 动态规划
	if len(prices) < 2 {
		return 0
	}
	dp, minPrice := make([]int, len(prices)), prices[0]
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}
	return dp[len(dp)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
