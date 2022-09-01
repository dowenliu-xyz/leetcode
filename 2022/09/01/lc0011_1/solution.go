package main

func maxArea(height []int) int {
	//return solution1(height)
	return solution2(height)
}

func solution1(height []int) int {
	// 暴力枚举 O(n^2) 超时
	var ans int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			ans = max(ans, (j-i)*min(height[i], height[j]))
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

func solution2(height []int) int {
	// 双指针
	if len(height) < 2 {
		return 0
	}
	l, r := 0, len(height)-1
	var ans int
	for l < r {
		ans = max(ans, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
