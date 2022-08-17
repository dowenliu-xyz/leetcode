package main

func generateParenthesis(n int) []string {
	//return solution1(n)
	return solution2(n)
}

func solution1(n int) []string {
	// 回溯，暴力枚举
	if n < 1 {
		return []string{}
	}
	var ans []string
	var isValid = func(s string) bool {
		left, right := 0, 0
		for _, c := range s {
			if c == '(' {
				left++
			} else {
				right++
			}
			if left < right {
				return false
			}
		}
		return left == right
	}
	var helper func(string)
	helper = func(s string) {
		if len(s) == n*2 {
			if isValid(s) {
				ans = append(ans, s)
			}
			return
		}
		helper(s + "(")
		helper(s + ")")
	}
	helper("")
	return ans
}

func solution2(n int) []string {
	// 剪枝
	if n < 1 {
		return []string{}
	}
	var ans []string
	var helper func(left, right int, s string)
	helper = func(left, right int, s string) {
		if left == n && right == n {
			ans = append(ans, s)
		}
		if left < n {
			helper(left+1, right, s+"(")
		}
		if right < left {
			helper(left, right+1, s+")")
		}
	}
	helper(0, 0, "")
	return ans
}
