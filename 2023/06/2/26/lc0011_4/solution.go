package lc0011

func maxArea(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	ans, left, right := 0, 0, n-1
	for left < right {
		ans = max(ans, (right-left)*min(height[left], height[right]))
		if height[left] <= height[right] {
			left++
		} else {
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
