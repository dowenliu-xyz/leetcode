package lc0416

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum&1 == 1 {
		return false
	}
	target := sum >> 1
	if max > target {
		return false
	}
	//dp := make([][]bool, n)
	//for i := range dp {
	//	dp[i] = make([]bool, target+1)
	//	dp[i][0] = true
	//}
	//dp[0][nums[0]] = true
	//for i := 1; i < n; i++ {
	//	for j := 1; j <= target; j++ {
	//		if j < nums[i] {
	//			dp[i][j] = dp[i-1][j]
	//		} else {
	//			dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
	//		}
	//	}
	//}
	//return dp[n-1][target]
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[target]
}
