package lc0045

func jump(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	step, leave, available := 0, 0, 0
	for i := 0; i < n; i++ {
		available = max(available, i+nums[i])
		if available >= n-1 {
			step++ // 可以到终点了，但不走是不会到的，还要再走最后一跳
			break
		}
		if i == leave { // 再前进就离开前一跳最大可达位置了，步数必须+1
			step++
			leave = available // 记录这一跳最大可达位置
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
