package main

import "math"

func maxSubArray(nums []int) int {
	//return solution1(nums)
	//return solution2(nums)
	//return solution3(nums)
	//return solution4(nums)
	return solution5(nums)
}

func solution1(nums []int) int {
	// 暴力枚举
	n := len(nums)
	ans := math.MinInt
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ans = max(ans, sum(nums, i, j))
		}
	}
	return ans
}

func sum(nums []int, i, j int) int {
	s := 0
	for k := i; k <= j; k++ {
		s += nums[k]
	}
	return s
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func solution2(nums []int) int {
	// dp
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

func solution3(nums []int) int {
	// dp 空间优化
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := nums[0]
	ans := dp
	for i := 1; i < n; i++ {
		dp = max(dp+nums[i], nums[i])
		ans = max(dp, ans)
	}
	return ans
}

func solution4(nums []int) int {
	// 贪心
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans, sum := math.MinInt, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		ans = max(ans, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

func solution5(nums []int) int {
	// 分治
	// 最大和子序列要么在中点一侧，要么横跨中点
	return get(nums, 0, len(nums)-1).maxSubSum
}

type Status struct {
	leftMaxSum  int
	rightMaxSum int
	totalSum    int
	maxSubSum   int
}

func get(nums []int, left, right int) Status {
	if left == right {
		return Status{
			leftMaxSum:  nums[left],
			rightMaxSum: nums[left],
			totalSum:    nums[left],
			maxSubSum:   nums[left],
		}
	}
	mid := left + (right-left)>>1
	s1, s2 := get(nums, left, mid), get(nums, mid+1, right)
	return pushUp(s1, s2)
}

func pushUp(left, right Status) Status {
	totalSum := left.totalSum + right.totalSum
	leftMaxSum := max(left.leftMaxSum, left.totalSum+right.leftMaxSum)
	rightMaxSum := max(right.rightMaxSum, left.rightMaxSum+right.totalSum)
	maxSubSum := max(max(left.maxSubSum, right.maxSubSum), left.rightMaxSum+right.leftMaxSum)
	return Status{
		leftMaxSum:  leftMaxSum,
		rightMaxSum: rightMaxSum,
		totalSum:    totalSum,
		maxSubSum:   maxSubSum,
	}
}
