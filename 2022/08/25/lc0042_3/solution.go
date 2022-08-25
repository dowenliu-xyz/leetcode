package main

func trap(height []int) int {
	//return solution1(height)
	//return solution2(height)
	//return solution3(height)
	return solution4(height)
}

func solution1(height []int) int {
	// 暴力枚举左右最大高度
	if len(height) < 3 {
		return 0
	}
	var ans int
	for i := 1; i < len(height)-1; i++ {
		left, right := height[i], height[i]
		for j := i - 1; j >= 0; j-- {
			left = max(left, height[j])
		}
		for j := i + 1; j < len(height); j++ {
			right = max(right, height[j])
		}
		ans += min(left, right) - height[i]
	}
	return ans
}

func solution2(height []int) int {
	// dp
	n := len(height)
	if n < 3 {
		return 0
	}
	leftMaxHeight, rightMaxHeight := make([]int, n), make([]int, n)
	leftMaxHeight[0], rightMaxHeight[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		leftMaxHeight[i] = max(leftMaxHeight[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMaxHeight[i] = max(rightMaxHeight[i+1], height[i])
	}
	var ans int
	for i := 1; i < n-1; i++ {
		ans += min(leftMaxHeight[i], rightMaxHeight[i]) - height[i]
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
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] <= height[i] {
			ground := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				leftSide := stack[len(stack)-1]
				width := i - leftSide - 1
				depth := min(height[leftSide], height[i]) - ground
				ans += width * depth
			}
		}
		stack = append(stack, i)
	}
	return ans
}

func solution4(height []int) int {
	// 双指针。dp方法的空间优化
	n := len(height)
	if n < 3 {
		return 0
	}
	var ans int
	leftMaxHeight, rightMaxHeight := 0, 0
	left, right := 0, n-1
	for left < right {
		if height[left] < height[right] { // 可能中间还会有更高的柱子，但现在可以保证当前相对另一边可能形成低洼
			if height[left] < leftMaxHeight { // 左右两边都比当前高，确定形成低洼
				ans += leftMaxHeight - height[left]
			} else { // 左边不比当前高，无法形成低洼
				leftMaxHeight = height[left]
			}
			left++ // 前进左指针
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
