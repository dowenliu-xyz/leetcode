package lc0322_1

import "math"

func coinChange(coins []int, amount int) int {
	//return solution1(coins, amount, make([]int, amount+1))
	return solution2(coins, amount)
}

func solution1(coins []int, remainAmount int, dp []int) int {
	if remainAmount < 0 {
		return -1
	}
	if remainAmount == 0 {
		return 0
	}
	if dp[remainAmount] != 0 {
		return dp[remainAmount]
	}
	min := math.MaxInt
	for i := range coins {
		res := solution1(coins, remainAmount-coins[i], dp)
		if res >= 0 && res < min {
			min = res + 1
		}
	}
	if min == math.MaxInt {
		dp[remainAmount] = -1
	} else {
		dp[remainAmount] = min
	}
	return dp[remainAmount]
}

func solution2(coins []int, amount int) int {
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
	dp[0] = 0 // 状态初始。找零 0元，无需任何硬币。
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i { // 可能能找开
				dp[i] = min(dp[i], dp[i-coins[j]]+1) // 看 i - coins[j] 是否能找开，其使用最少硬币数是多少，+ 1 即找开 i 的最少硬币数.
			}
		}
	}
	if dp[amount] <= amount { // 一种极端情况，只有 1元 硬币时，dp[amount] == amount 成立
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
