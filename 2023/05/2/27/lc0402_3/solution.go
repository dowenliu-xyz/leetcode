package lc0402

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}
