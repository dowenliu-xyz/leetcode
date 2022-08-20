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
	// 暴力枚举 超时
	ans := math.MinInt
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			ans = max(ans, sum(nums[i:j+1]))
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func sum(nums []int) int {
	var ans int
	for _, num := range nums {
		ans += num
	}
	return ans
}

func solution2(nums []int) int {
	// dp
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		//*/
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		/*/
		dp[i] = max(dp[i-1] + nums[i], nums[i])
		//*/
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
	dp, ans := nums[0], nums[0]
	for i := 1; i < n; i++ {
		dp = max(dp+nums[i], nums[i])
		ans = max(ans, dp)
	}
	return ans
}

func solution4(nums []int) int {
	// 贪心
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans, sum := nums[0], 0
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
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

type Status struct {
	lSum, rSum, mSum, iSum int
}
