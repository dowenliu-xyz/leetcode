package lc0035_2

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l++
		} else {
			r--
		}
	}
	return l
}
