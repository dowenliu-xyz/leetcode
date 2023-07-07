package lc0032

func longestValidParentheses(s string) int {
	//return solution1(s)
	//return solution2(s)
	return solution3(s)
}

func solution1(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	ans := 0
	dp := make([]int, n) // dp[i] 表示以 i 为右边界的最长有效括号子串的长度
	for i := 1; i < n; i++ {
		if s[i] == '(' { // 不能作为右边界
			continue
		}
		left := i - dp[i-1] - 1         // 当前位置对应左边界
		if left < 0 || s[left] == ')' { // 左边界无效
			continue
		}
		if left >= 1 { // 检查包围后再左侧是否可能连接另一个有效括号串
			// dp[left-1] 要么是 0 （即没有连接有效括号串）要么大于 0 （即连接一个有效括号串）
			dp[i] = dp[i-1] + 2 + dp[left-1]
		} else { // left 已经是 s 开头
			dp[i] = dp[i-1] + 2
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func solution2(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	stack := make([]int, 0, n+1) // 栈底始终保持最后一个无法配对的右括号的位置
	stack = append(stack, -1)
	ans := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' { // 左括号直接入栈
			stack = append(stack, i)
			continue
		}
		// s[i] == ')'
		stack = stack[:len(stack)-1]
		if len(stack) == 0 { // 出栈的是栈底
			stack = append(stack, i) // 当前是未匹配的右括号，替换入栈底
		} else {
			ans = max(ans, i-stack[len(stack)-1]) // 计算当前有效串长，尝试更新结果
		}
	}
	return ans
}

func solution3(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	left, right, ans := 0, 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			ans = max(ans, left+right)
		} else if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			ans = max(ans, left+right)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
