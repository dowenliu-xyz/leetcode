package main

func canJump(nums []int) bool {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) bool {
	// 反向扫描
	n := len(nums)
	if n < 2 {
		return true
	}
	target := n - 1
	for target > 0 {
		tmp := target
		for i := 0; i < target; i++ {
			if i+nums[i] >= target {
				target = i
				break
			}
		}
		if tmp == target {
			return false
		}
	}
	return true
}

func solution2(nums []int) bool {
	// 贪心
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if i > reachable {
			return false
		}
		reachable = max(reachable, i+nums[i])
		if reachable >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
