package lc0402_1

func removeKdigits(num string, k int) string {
	// case: 1234567890 移除 9 位 , 结果 0，明显是单调栈解法
	stack := make([]byte, 0, len(num))
	for i := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			// 仍有移除额度，且前一位比当前位大
			stack = stack[:len(stack)-1]
			k-- // 减少要移除位数额度
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k] // 可能还有 k 位要移除，移除尾部的
	// 清除开头的 '0'
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	if len(stack) == 0 {
		// 特殊情况：所有数字都被移除了
		return "0"
	}
	return string(stack)
}
