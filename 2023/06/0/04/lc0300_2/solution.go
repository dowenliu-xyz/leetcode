package lc0300

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0, len(nums))
	for _, num := range nums {
		if len(dp) > 0 && dp[len(dp)-1] >= num {
			l, r := 0, len(dp)-1
			for l <= r {
				m := l + (r-l)>>1
				if num <= dp[m] {
					r = m - 1
				} else {
					l = m + 1
				}
			}
			dp[l] = num
		} else {
			dp = append(dp, num)
		}
	}
	return len(dp)
}
