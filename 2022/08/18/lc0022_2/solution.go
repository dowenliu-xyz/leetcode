package main

func generateParenthesis(n int) []string {
	//return solution1(n)
	return solution2(n)
}

func solution1(n int) []string {
	// 傻递归
	if n < 1 {
		return []string{}
	}
	var ans []string
	isValid := func(s []byte) bool {
		l, r := 0, 0
		for _, b := range s {
			if b == '(' {
				l++
			} else {
				r++
			}
			if l < r {
				return false
			}
		}
		return l == r
	}
	var helper func(accum []byte)
	helper = func(accum []byte) {
		if len(accum) == n*2 {
			if isValid(accum) {
				ans = append(ans, string(accum))
			}
			return
		}
		helper(append(accum, '('))
		helper(append(accum, ')'))
	}
	helper(nil)
	return ans
}

func solution2(n int) []string {
	// 剪枝
	if n < 1 {
		return []string{}
	}
	var ans []string
	var helper func(l, r int, accum []byte)
	helper = func(l, r int, accum []byte) {
		if l == n && r == n {
			ans = append(ans, string(accum))
			return
		}
		if l < n {
			helper(l+1, r, append(accum, '('))
		}
		if l > r {
			helper(l, r+1, append(accum, ')'))
		}
	}
	helper(0, 0, nil)
	return ans
}
