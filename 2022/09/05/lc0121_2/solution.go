package main

func maxProfit(prices []int) int {
	// dp
	// dp[i] 第i天最大收益 i>0
	// dp[0] = 0 因为当天只能买入。即使允许当天买入当天卖出收益也为0
	// 当i>0时，假设已经持有股票，则有：
	// 1）如果当日价格不比前一天高，则当日最大收益与前一天最大收益相同
	// 2）如果当日价格比前一天高，则今日卖出收益较高
	// 第i天（i>0)最大收益为当天价格-第0到i-1天之前最低价格
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp, minPrice := make([]int, n), prices[0]
	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}
	return dp[n-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
