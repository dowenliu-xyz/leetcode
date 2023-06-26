package lc0045

func jump(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	ans, leave, reachable := 0, 0, 0
	for i := 0; i < n-1; i++ {
		reachable = max(reachable, i+nums[i])
		if i == leave {
			ans++
			leave = reachable
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
