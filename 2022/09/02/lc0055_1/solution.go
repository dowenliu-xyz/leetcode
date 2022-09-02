package main

func canJump(nums []int) bool {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) bool {
	// 反向查找
	if len(nums) < 2 {
		return true
	}
	target := len(nums) - 1
	for {
		p := target
		for i := 0; i < target; i++ {
			if i+nums[i] >= target {
				if i == 0 {
					return true
				}
				target = i
				break
			}
		}
		if p == target {
			return false
		}
	}
}

func solution2(nums []int) bool {
	// 贪心
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if i > reachable {
			return false
		}
		reachable = max(reachable, i+nums[i])
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
