package lc0033

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] {
			// [l:m] 为递增区间
			if nums[l] <= target && target < nums[m] {
				// target 在该递增区间内
				r = m - 1
			} else {
				// target 只能在 [m+1:r] 这个非递增区间内。缩小范围继续排查
				l = m + 1
			}
		} else {
			// nums[m:r+1] 为递增区间
			if nums[m] < target && target <= nums[r] {
				// target 在该递增区间内
				l = m + 1
			} else {
				// target 只能在 [l:m] 这个非递增区间内。缩小范围继续排查
				r = m - 1
			}
		}
	}
	return -1
}
