package lc0322_2

import "math"

func coinChange(coins []int, amount int) int {
	//return solution1(coins, amount, make([]int, amount+1))
	return solution2(coins, amount)
}

func solution1(coins []int, rem int, mem []int) int {
	// 记忆化搜索
	if rem < 0 {
		// 找不开
		return -1
	}
	if rem == 0 {
		// 找 0 元，无需任何硬币
		return 0
	}
	if mem[rem] != 0 {
		// 计算过 rem 金额的最少找零硬币数
		return mem[rem]
	}
	// 遍历各币值后，找使用一枚硬币后，是否能找开、能找开的话最少找零硬币数
	min := math.MaxInt
	for i := range coins {
		n := solution1(coins, rem-coins[i], mem)
		if n >= 0 && n < min {
			min = n
		}
	}
	if min == math.MaxInt {
		// 找不开
		mem[rem] = -1
	} else {
		// 找到的硬币数再加上当前使用了一枚硬币
		mem[rem] = min + 1
	}
	return mem[rem]
}

func solution2(coins []int, amount int) int {
	// dp
	// 先检查特殊值
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = amount + 1 // 使用所有 dp[amount] <= amount 不成立
	for i := 1; i < len(dp); i = i << 1 {
		copy(dp[i:], dp[:i])
	}
	dp[0] = 0 // 找零 0 元只有一种方案，一个硬币都不出
	// 推算各金额找零方案
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				// 有可能找开
				dp[i] = min(dp[i], dp[i-coin]+1) // 要么找不开，要么当前减去当前币值能找开
			}
		}
	}
	if dp[amount] <= amount {
		return dp[amount]
	}
	return -1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
