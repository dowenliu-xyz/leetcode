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
	// 暴力枚举。应该会超时
	ans := math.MinInt
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			ans = max(ans, sumInts(nums[i:j+1]))
		}
	}
	return ans
}

func sumInts(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans += num
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
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
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

func solution3(nums []int) int {
	// dp 空间优化
	if len(nums) == 0 {
		return 0
	}
	dp, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp = max(dp+nums[i], nums[i])
		ans = max(ans, dp)
	}
	return ans
}

func solution4(nums []int) int {
	// 贪心
	if len(nums) == 0 {
		return 0
	}
	ans, sum := math.MinInt, 0
	for i := 0; i < len(nums); i++ {
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
	// 区间最大子序列和子序列要么在中间分隔线左侧，要么在分隔线右侧，要么横跨分隔线。只能是三种情况之一。
	return get(nums, 0, len(nums)-1).maxSubSum
}

func get(nums []int, left, right int) Status {
	if left == right {
		// 最基础的情况，区间中只有一个元素。
		return Status{
			beginMaxSubSum: nums[left],
			endMaxSubSum:   nums[left],
			maxSubSum:      nums[left],
			allSum:         nums[left],
		}
	}
	mid := left + (right-left)>>1
	leftStatus := get(nums, left, mid)
	rightStatus := get(nums, mid+1, right)
	return pushUp(leftStatus, rightStatus)
}

func pushUp(left, right Status) Status {
	maxSubSum := max(
		// 区间最大子序和子序列，要么在分隔线一侧，即 左侧子区间 maxSubSum 或 右侧子区间 maxSubSum
		max(left.maxSubSum, right.maxSubSum),
		// 要么横跨分隔线。
		// 此时，子序列由"左侧子区间 endMaxSubArray"和"右侧子区间 beginMaxSubArray"构成。
		left.endMaxSubSum+right.beginMaxSubSum,
		// 三者取大
	)
	// 开始节点为开头的最大子序和要么是"左侧区间的 beginMaxSubSum，要么是"左侧区间 allSum + 右侧区间 beginMaxSubSum"
	// 二者取大
	beginMaxSubSum := max(left.beginMaxSubSum, left.allSum+right.beginMaxSubSum)
	// 同理，可得 结尾节点为结尾的最大子序列和
	endMaxSubSum := max(right.endMaxSubSum, left.endMaxSubSum+right.allSum)
	allSum := left.allSum + right.allSum
	return Status{
		beginMaxSubSum: beginMaxSubSum,
		endMaxSubSum:   endMaxSubSum,
		maxSubSum:      maxSubSum,
		allSum:         allSum,
	}
}

type Status struct {
	// 以区间开头节点为起点的最大子序列和
	beginMaxSubSum int
	// 以区间结尾节点为结尾的最大子序列和
	endMaxSubSum int
	// 区间内最大子序列和
	maxSubSum int
	// 区间和
	allSum int
}
