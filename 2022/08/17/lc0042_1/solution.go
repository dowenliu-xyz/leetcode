package main

func trap(height []int) int {
	//return solution1(height)
	//return solution2(height)
	//return solution3(height)
	return solution4(height)
}

func solution1(height []int) int {
	// 暴力
	var ans int
	if len(height) < 3 {
		return 0
	}
	for i := 1; i < len(height)-1; i++ {
		leftMaxHeight, rightMaxHeight := 0, 0
		for j := 0; j < i; j++ {
			leftMaxHeight = max(height[j], leftMaxHeight)
		}
		for j := i + 1; j < len(height); j++ {
			rightMaxHeight = max(height[j], rightMaxHeight)
		}
		maxHeight := min(leftMaxHeight, rightMaxHeight)
		ans += max(maxHeight-height[i], 0)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func solution2(height []int) int {
	// dp
	if len(height) < 3 {
		return 0
	}
	var ans int
	leftHeight, rightHeight := make([]int, len(height)), make([]int, len(height))
	for i := 1; i < len(height); i++ {
		leftHeight[i] = max(height[i-1], leftHeight[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightHeight[i] = max(height[i+1], rightHeight[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		ans += max(0, min(leftHeight[i], rightHeight[i])-height[i])
	}
	return ans
}

func solution3(height []int) int {
	// 栈
	if len(height) < 3 {
		return 0
	}
	var ans int
	var stack []int
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			low := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				width := i - left - 1
				depth := min(height[left], height[i]) - low
				ans += width * depth
			}
		}
		stack = append(stack, i)
	}
	return ans
}

func solution4(height []int) int {
	// 双指针
	if len(height) < 3 {
		return 0
	}
	var ans int
	leftMax, rightMax := height[0], height[len(height)-1]
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
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
