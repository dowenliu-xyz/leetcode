package lc0070_4

func climbStairs(n int) int {
	//return solution1(n)
	return solution2(n)
}

func solution1(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n)
	dp[0], dp[1] = 1, 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func solution2(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	p, q := 1, 2
	for i := 2; i < n; i++ {
		p, q = q, p+q
	}
	return q
}
