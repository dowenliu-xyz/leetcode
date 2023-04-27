package lc0402_2

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num))
	for i := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
