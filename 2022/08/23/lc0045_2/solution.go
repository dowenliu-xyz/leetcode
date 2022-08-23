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
	target := len(nums) - 1
	steps := 0
	for target > 0 {
		for i := 0; i < target; i++ {
			if i+nums[i] >= target {
				target = i
				steps++
				break
			}
		}
	}
	return steps
}

func solution2(nums []int) int {
	// 贪心
	end, maxPosition, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; /*当可到达终点时就不用再跳了*/ i++ {
		maxPosition = max(maxPosition, i+nums[i]) // 更新下一跳可达区间最大范围
		if i == end {                             // 当前区间范围已经检查完成
			// 跳到下一可达范围
			end = maxPosition
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
