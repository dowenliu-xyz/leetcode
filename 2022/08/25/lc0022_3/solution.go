package main

func generateParenthesis(n int) []string {
	//return solution1(n)
	return solution2(n)
}

func solution1(n int) []string {
	// 回溯，傻递归
	if n < 1 {
		return []string{}
	}
	var ans []string
	helper(n*2, 0, nil, &ans)
	return ans
}

func helper(maxChars, curChars int, a []byte, ans *[]string) {
	if curChars == maxChars {
		if isValid(a) {
			*ans = append(*ans, string(a))
		}
		return
	}
	helper(maxChars, curChars+1, append(a, '('), ans)
	helper(maxChars, curChars+1, append(a, ')'), ans)
}

func isValid(a []byte) bool {
	if len(a) == 0 {
		return true
	}
	left, right := 0, 0
	for _, b := range a {
		if b == '(' {
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

func solution2(n int) []string {
	// 回溯，剪枝
	if n < 1 {
		return []string{}
	}
	var ans []string
	helper2(n, 0, 0, nil, &ans)
	return ans
}

func helper2(n, left, right int, a []byte, ans *[]string) {
	if left == n && right == n {
		*ans = append(*ans, string(a))
		return
	}
	if left < n {
		helper2(n, left+1, right, append(a, '('), ans)
	}
	if left > right {
		helper2(n, left, right+1, append(a, ')'), ans)
	}
}
