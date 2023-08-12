package lc0055

func canJump(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}
	cur, reachable, last := 0, 0, n-1
	for {
		reachable = max(reachable, cur+nums[cur])
		if reachable >= last {
			return true
		}
		cur++
		if cur > reachable {
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
