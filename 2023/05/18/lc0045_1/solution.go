package lc0045_1

func jump(nums []int) int {
	leave, reachable, ans := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		reachable = max(reachable, i+nums[i])
		if i == leave {
			leave = reachable
			ans++
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
