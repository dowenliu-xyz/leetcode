package main

func trap(height []int) int {
	//return solution1(height)
	//return solution2(height)
	//return solution3(height)
	return solution4(height)
}

func solution1(height []int) int {
	// 暴力枚举积水深度
	length := len(height)
	if length < 3 {
		// 宽度不够，不足以形成低洼
		return 0
	}
	var ans int
	for i := 1; i < length-1; i++ {
		leftHeight, rightHeight := height[0], height[length-1]
		for j := 0; j < i; j++ {
			leftHeight = max(leftHeight, height[j])
		}
		for j := length - 2; j > i; j-- {
			rightHeight = max(rightHeight, height[j])
		}
		depth := max(0, min(leftHeight, rightHeight)-height[i])
		ans += depth
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
	n := len(height)
	if n < 3 {
		return 0
	}
	var ans int
	leftHeight, rightHeight := make([]int, n), make([]int, n)
	leftHeight[0], rightHeight[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		leftHeight[i] = max(leftHeight[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rightHeight[i] = max(rightHeight[i+1], height[i+1])
	}
	for i := 1; i < n-1; i++ {
		depth := max(0, min(leftHeight[i], rightHeight[i])-height[i])
		ans += depth
	}
	return ans
}

func solution3(height []int) int {
	// 单调栈
	n := len(height)
	if n < 3 {
		return 0
	}
	var ans int
	var stack []int // 栈元素为下标
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			ground := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				leftSide := stack[len(stack)-1]
				width := i - leftSide - 1
				leftHeight := height[leftSide]
				depth := max(0, min(leftHeight, height[i])-ground)
				delta := width * depth
				ans += delta
			}
		}
		stack = append(stack, i)
	}
	return ans
}

func solution4(height []int) int {
	// 双指针
	n := len(height)
	if n < 3 {
		return 0
	}
	var ans int
	leftMaxHeight, rightMaxHeight := 0, 0
	left, right := 0, n-1
	for left < right {
		if height[left] < height[right] {
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
