package main

func romanToInt(s string) int {
	//return solution1(s)
	return solution2(s)
}

func solution1(s string) int {
	// 反向扫描
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	if len(s) == 0 {
		return 0
	}
	ans := values[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		cur, next := s[i], s[i+1]
		value := values[cur]
		if cur == 'I' && (next == 'V' || next == 'X') {
			ans -= value
		} else if cur == 'X' && (next == 'L' || next == 'C') {
			ans -= value
		} else if cur == 'C' && (next == 'D' || next == 'M') {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}

func solution2(s string) int {
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	ans := 0
	for i := 0; i < len(s)-1; i++ {
		value := values[s[i]]
		if value < values[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	ans += values[s[len(s)-1]]
	return ans
}
