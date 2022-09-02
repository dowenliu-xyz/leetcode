package main

func maxArea(height []int) int {
	// 双指针
	if len(height) < 2 {
		return 0
	}
	l, r, ans := 0, len(height)-1, 0
	for l < r {
		ans = max(ans, (r-l)*min(height[l], height[r]))
		if height[l] <= height[r] {
			l++
		} else {
			r--
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
