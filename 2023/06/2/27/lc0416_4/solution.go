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
	half := sum >> 1
	if max == half {
		return true
	}
	if max > half {
		return false
	}
	dp := make([]bool, half+1) // dp[i] 表示元素和能否是 i
	dp[0] = true
	for _, num := range nums {
		// 每个元素只能用一次
		for i := half; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[half]
}
