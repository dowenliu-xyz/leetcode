package lc0020

func isValid(s string) bool {
	pairs := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == pairs[s[i]] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}
