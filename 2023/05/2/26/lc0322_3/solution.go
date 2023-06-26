package lc0322

import "math"

func coinChange(coins []int, amount int) int {
	//return solution1(coins, amount, make([]int, amount+1))
	return solution2(coins, amount)
}

func solution1(coins []int, amount int, mem []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if mem[amount] != 0 {
		return mem[amount]
	}
	minUse := math.MaxInt
	for _, coin := range coins {
		use := solution1(coins, amount-coin, mem)
		if use == -1 {
			continue
		}
		if use < minUse {
			minUse = use
		}
	}
	if minUse == math.MaxInt {
		mem[amount] = -1
	} else {
		mem[amount] = minUse + 1
	}
	return mem[amount]
}

func solution2(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = amount + 1
	for i := 1; i < len(dp); i = i << 1 {
		copy(dp[i:], dp[:i])
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
	if dp[amount] <= amount {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
