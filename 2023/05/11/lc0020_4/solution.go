package lc0020_4

func isValid(s string) bool {
	pairs := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if p, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != p {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
