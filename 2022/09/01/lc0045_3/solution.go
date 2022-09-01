package main

func jump(nums []int) int {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) int {
	// 反向扫描
	if len(nums) < 2 {
		return 0
	}
	ans, target := 0, len(nums)-1
	for target > 0 {
		for i := 0; i < target; i++ {
			if i+nums[i] >= target {
				target = i
				ans++
				break
			}
		}
	}
	return ans
}

func solution2(nums []int) int {
	// 贪心
	leave, reachable, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		reachable = max(reachable, i+nums[i])
		if i == leave {
			leave = reachable
			steps++
		}
	}
	return steps
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
