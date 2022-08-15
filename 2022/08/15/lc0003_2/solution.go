package main

func lengthOfLongestSubstring(s string) int {
	//return solution1(s)
	return solution2(s)
}

func solution1(s string) int {
	// 暴力枚举。O(n^3) 超时
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isUnique(s[i : j+1]) {
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}

func isUnique(s string) bool {
	seen := map[rune]bool{}
	for _, b := range s {
		if seen[b] {
			return false
		}
		seen[b] = true
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func solution2(s string) int {
	// 滑动窗口
	window := map[byte]bool{}
	left, right, ans := 0, 0, 0
	for right < len(s) {
		for window[s[right]] {
			if len(window) > ans {
				ans = len(window)
			}
			delete(window, s[left])
			left++
		}
		window[s[right]] = true
		right++
	}
	if len(window) > ans {
		ans = len(window)
	}
	return ans
}
