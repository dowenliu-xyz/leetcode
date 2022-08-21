package main

func jump(nums []int) int {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) int {
	// 反向查找出发位置
	position := len(nums) - 1 // 目的地，初始为终点
	steps := 0
	for position > 0 { // 循环条件，目的地不是起点
		for i := 0; i < position; i++ { // 循环条件，在目的地前出发
			if i+nums[i] >= position { // 从当前点可以到达目的地，找到最前的出发点。
				steps++
				position = i // 之后开始以该点为目的地找最早出发点
				break
			}
		}
	}
	return steps
}

func solution2(nums []int) int {
	// 贪心。正向查找可到达的最大位置
	end := 0         // 上一跳可达位置
	maxPosition := 0 // 下一跳最远可达位置
	steps := 0
	for i := 0; i < len(nums)-1; /*如果这里是 len(nums)，就相当于在终点原地在跳一下，计数会超*/ i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
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
