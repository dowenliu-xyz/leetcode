package main

func lengthOfLongestSubstring(s string) int {
	//return solution1(s)
	return solution2(s)
}

func solution1(s string) int {
	// 暴力枚举
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isUnique(s[i:j+1]) && ans < j-i+1 {
				ans = j - i + 1
			}
		}
	}
	return ans
}

func isUnique(s string) bool {
	seen := map[rune]bool{}
	for _, r := range s {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func solution2(s string) int {
	// 滑动窗口
	ans := 0
	window := map[byte]bool{}
	left, right := 0, 0
	for right < len(s) {
		for window[s[right]] { // 有重复
			ans = max(ans, len(window))
			delete(window, s[left])
			left++
		}
		window[s[right]] = true
		right++
	}
	if len(window) > 0 { // 处理最后一段无重复子串
		ans = max(ans, len(window))
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
