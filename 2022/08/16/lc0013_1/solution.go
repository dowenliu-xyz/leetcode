package main

func romanToInt(s string) int {
	//return solution1(s)
	return solution2(s)
}

func solution1(s string) int {
	// 反向遍历 + 模拟
	table := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var ans int
	var r byte
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		v := table[c]
		if c == 'I' && (r == 'V' || r == 'X') {
			ans -= v
		} else if c == 'X' && (r == 'L' || r == 'C') {
			ans -= v
		} else if c == 'C' && (r == 'D' || r == 'M') {
			ans -= v
		} else {
			ans += v
		}
		r = c
	}
	return ans
}

func solution2(s string) int {
	// 因题目保证输出总是合法罗马数字，这里忽略非法组合，可以简化写法
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	for i := range s {
		v := values[s[i]]
		if i < len(s)-1 && v < values[s[i+1]] {
			ans -= v
		} else {
			ans += v
		}
	}
	return ans
}
