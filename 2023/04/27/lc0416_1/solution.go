package lc0416_1

import "math"

func canPartition(nums []int) bool {
	n := len(nums)
	if n <= 1 { // 特殊情况，只有一个元素，一定不能分成两个集合
		return false
	}
	// 计算所有元素和 和 最大元素值
	sum, max := 0, math.MinInt
	for i := 0; i < n; i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	if sum%2 == 1 { // 和为奇数，无法均分为两个整数集合
		return false
	}
	target := sum / 2
	if max > target { // 最大元素大于 target，剩下的一定达不到 target
		return false
	}
	if max == target { // 最大元素正好一半
		return true
	}
	// dp 数组，每行表示考察第 n 个元素后可构成的元素和。每行 target + 1 个元素。首列是辅助列，表示是否元素和为 0。
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// 处理首行，只有两种情况，选择或不选 nums[0]
	dp[0][0] = true          // 不选，和为 0
	dp[0][nums[0]] = true    // 选择，和为 nums[0]
	for i := 1; i < n; i++ { // 从第二个元素开始考察
		for j := 0; j <= target; j++ { // 判断是否能实现元素和
			dp[i][j] = dp[i][j] || dp[i-1][j] // 不选当前元素则检查 dp[i-1][j]
			if j >= nums[i] {                 // 当前元素小于 j 时，检查 dp[i-1][j-nums[i]]
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][target] // 最后一行的 target 列即结果
}
