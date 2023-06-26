package lc0055

func canJump(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}
	cur, available, last := 0, 0, n-1
	for {
		available = max(available, cur+nums[cur])
		if available >= last {
			return true
		}
		cur++
		if cur > available {
			return false
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
