package lc0055_3

func canJump(nums []int) bool {
	reachable := 0
	for i, num := range nums {
		if i > reachable {
			return false
		}
		reachable = max(reachable, i+num)
		if reachable >= len(nums)-1 {
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
