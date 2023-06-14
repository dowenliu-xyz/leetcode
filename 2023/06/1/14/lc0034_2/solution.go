package lc0034

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if target <= nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if l >= len(nums) || nums[l] != target {
		return ans
	}
	ans[0] = l
	r = len(nums) - 1
	for l <= r {
		m := l + (r-l)>>1
		if target >= nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	ans[1] = r
	return ans
}
