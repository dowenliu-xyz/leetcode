package main

func climbStairs(n int) int {
	//return solution1(n)
	//return solution2(n)
	//return solution3(n)
	return solution4(n)
}

func solution1(n int) int {
	// 傻递归
	if n < 2 {
		return 1
	}
	return solution1(n-1) + solution1(n-2)
}

func solution2(n int) int {
	// dp 加速
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	var helper func(n int) int
	helper = func(n int) int {
		if dp[n] == 0 {
			dp[n] = helper(n-2) + helper(n-1)
		}
		return dp[n]
	}
	return helper(n)
}

func solution3(n int) int {
	// dp 迭代形式
	if n < 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func solution4(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return q
}
