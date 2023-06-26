package lc0042

func trap(height []int) int {
	//return solution1(height)
	return solution2(height)
}

func solution1(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	dp := make([][]int, n) // dp[i] 表示在i处左右柱子最大高度
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[n-1][1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		dp[i][1] = max(dp[i+1][1], height[i])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		ans += max(0, min(dp[i][0], dp[i][1])-height[i])
	}
	return ans
}

func solution2(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	ans, left, right, leftMaxHeight, rightMaxHeight := 0, 0, n-1, 0, 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] < leftMaxHeight {
				ans += leftMaxHeight - height[left]
			} else {
				leftMaxHeight = height[left]
			}
			left++
		} else {
			if height[right] < rightMaxHeight {
				ans += rightMaxHeight - height[right]
			} else {
				rightMaxHeight = height[right]
			}
			right--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
