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
		// s[i] 必然是 ')'
		left := i - dp[i-1] - 1         // 使用公式直接计算当前位置对应左边界
		if left < 0 || s[left] == ')' { // 计算出的左边界无效
			continue
		}
		if left >= 1 { // 左侧还有可能连接着一个有效括号串
			// dp[left-1] 可能是 0 （左边没法扩展了），也可能不是 0 （可以向左再连接）
			dp[i] = dp[i-1] + 2 + dp[left-1]
		} else {
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
	// stack 用于判断是否有效括号子串，元素为索引位置，其底始终为上一个未能配对的括号位置
	stack := make([]int, 0, n+1)
	stack = append(stack, -1)
	ans := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
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
