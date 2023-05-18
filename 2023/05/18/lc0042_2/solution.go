package lc0042_2

func trap(height []int) int {
	//return solution1(height)
	//return solution2(height)
	return solution3(height)
}

func solution1(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0], rightMax[n-1] = 0, 0
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += max(0, min(leftMax[i], rightMax[i])-height[i])
	}
	return ans
}

func solution2(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	ans := 0
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			low := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				minHigh := min(height[stack[len(stack)-1]], height[i])
				depth := max(0, minHigh-low)
				width := i - stack[len(stack)-1] - 1
				ans += depth * width
			}
		}
		stack = append(stack, i)
	}
	return ans
}

func solution3(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	ans, leftMax, rightMax, left, right := 0, 0, 0, 0, n-1
	for left < right {
		if height[left] <= height[right] {
			if height[left] < leftMax {
				ans += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				ans += rightMax - height[right]
			} else {
				rightMax = height[right]
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
