package lc0055_2

func canJump(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
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
