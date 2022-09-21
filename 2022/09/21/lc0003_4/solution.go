package main

func lengthOfLongestSubstring(s string) int {
	//return solution1(s)
	return solution2(s)
}

func solution1(s string) int {
	// 暴力
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if j-i+1 > ans && isUnique(s, i, j) {
				ans = j - i + 1
			}
		}
	}
	return ans
}

func isUnique(s string, l, r int) bool {
	seen := map[byte]struct{}{}
	for i := l; i <= r; i++ {
		if _, ok := seen[s[i]]; ok {
			return false
		}
		seen[s[i]] = struct{}{}
	}
	return true
}

func solution2(s string) int {
	// 滑动窗口
	w := map[byte]struct{}{}
	ans := 0
	l := 0
	for i := 0; i < len(s); i++ {
		for {
			if _, ok := w[s[i]]; ok {
				delete(w, s[l])
				l++
			} else {
				break
			}
		}
		w[s[i]] = struct{}{}
		ans = max(ans, len(w))
	}
	ans = max(ans, len(w))
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
