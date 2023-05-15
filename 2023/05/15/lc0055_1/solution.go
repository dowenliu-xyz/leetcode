package lc0055_1

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	reachable, target := 0, n-1
	for i, num := range nums {
		if i > reachable {
			return false
		}
		reachable = max(reachable, i+num)
		if reachable >= target {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
