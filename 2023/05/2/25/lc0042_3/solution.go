package lc0042

func trap(height []int) int {
	//return solution1(height)
	//return solution2(height)
	return solution3(height)
}

func solution1(height []int) int {
	// dp
	n := len(height)
	if n < 3 {
		return 0
	}
	leftMaxHeight, rightMaxHeight := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		leftMaxHeight[i] = max(leftMaxHeight[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rightMaxHeight[i] = max(rightMaxHeight[i+1], height[i+1])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += max(0, min(leftMaxHeight[i], rightMaxHeight[i])-height[i])
	}
	return ans
}

func solution2(height []int) int {
	// 单调栈
	n := len(height)
	if n < 3 {
		return 0
	}
	ans := 0
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] <= height[i] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				width := i - stack[len(stack)-1] - 1
				depth := max(0, min(height[stack[len(stack)-1]], height[i])-bottom)
				ans += width * depth
			}
		}
		stack = append(stack, i)
	}
	return ans
}

func solution3(height []int) int {
	// 双指针
	n := len(height)
	if n < 3 {
		return 0
	}
	leftMaxHeight, rightMaxHeight, left, right := 0, 0, 0, n-1
	ans := 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] < leftMaxHeight {
				ans += leftMaxHeight - height[left]
			} else {
				leftMaxHeight = height[left]
			}
			left++
		} else {
			if height[right] < rightMaxHeight {
				ans += rightMaxHeight - height[right]
			} else {
				rightMaxHeight = height[right]
			}
			right--
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
