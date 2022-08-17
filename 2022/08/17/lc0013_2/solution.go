package main

func romanToInt(s string) int {
	//return solution1(s)
	return solution2(s)
}

func solution1(s string) int {
	// 模拟
	values := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		value := values[c]
		if c == 'I' {
			if i+1 < len(s) {
				if s[i+1] == 'V' || s[i+1] == 'X' {
					ans -= value
				} else {
					ans += value
				}
			} else {
				ans += value
			}
		} else if c == 'X' {
			if i+1 < len(s) {
				if s[i+1] == 'L' || s[i+1] == 'C' {
					ans -= value
				} else {
					ans += value
				}
			} else {
				ans += value
			}
		} else if c == 'C' {
			if i+1 < len(s) {
				if s[i+1] == 'D' || s[i+1] == 'M' {
					ans -= value
				} else {
					ans += value
				}
			} else {
				ans += value
			}
		} else {
			ans += value
		}
	}
	return ans
}

func solution2(s string) int {
	// 利用题目约束化简
	// 因题目保证输入为合法罗马数字，可将模拟方法化简
	values := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	ans := 0
	for i := range s {
		v := values[s[i]]
		if i+1 < len(s) && v < values[s[i+1]] {
			ans -= v
		} else {
			ans += v
		}
	}
	return ans
}
