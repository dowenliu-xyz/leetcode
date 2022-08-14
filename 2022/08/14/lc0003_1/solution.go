package main

func lengthOfLongestSubstring(s string) int {
	//return solution1(s)
	return solution2(s)
}

func solution1(s string) int {
	// 暴力枚举，使用 map 判重
	maxLen := 0
	for i := 0; i < len(s); i++ {
		seen := map[byte]bool{}
		for j := i; j <= len(s); j++ {
			if j == len(s) || seen[s[j]] {
				if j-i > maxLen {
					maxLen = j - i
				}
				break
			}
			seen[s[j]] = true
		}
	}
	return maxLen
}

func solution2(s string) int {
	// 滑动窗口
	window, ans := map[byte]bool{}, 0
	begin, end := 0, 0
	for end < len(s) {
		for window[s[end]] {
			if len(window) > ans {
				ans = len(window)
			}
			delete(window, s[begin])
			begin++
		}
		window[s[end]] = true
		end++
	}
	if len(window) > ans {
		ans = len(window)
	}
	return ans
}
