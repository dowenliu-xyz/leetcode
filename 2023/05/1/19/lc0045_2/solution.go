package lc0045_2

func jump(nums []int) int {
	// 贪心
	leave, reachable, steps := 0, 0, 0
	// 注意 i < len(nums)-1，而不是 i < len(nums)。因为不用从 nums[len(nums)-1] 离开
	for i := 0; i < len(nums)-1; i++ {
		reachable = max(reachable, i+nums[i])
		if i == leave {
			leave = reachable
			steps++
		}
	}
	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
