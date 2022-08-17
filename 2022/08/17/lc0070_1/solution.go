package main

func climbStairs(n int) int {
	//return solution1(n)
	//return solution2(n)
	return solution3(n)
}

func solution1(n int) int {
	// 递归
	if n < 2 {
		return 1
	}
	return solution1(n-1) + solution1(n-2)
}

func solution2(n int) int {
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

func solution3(n int) int {
	if n < 2 {
		return 1
	}
	p, q := 1, 1
	for i := 2; i <= n; i++ {
		p, q = q, p+q
	}
	return q
}
