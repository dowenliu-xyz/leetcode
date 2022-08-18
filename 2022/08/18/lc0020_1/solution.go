package main

func isValid(s string) bool {
	table := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	var stack []byte
	for i := range s {
		c := s[i]
		if v, ok := table[c]; ok {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 || c != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

// 还有一种比较少想到的方法
// 不断依次替换三种括号对为空白，直到字符串全部换成空白，如果最终能换成空白说明成对；
// 每轮次三种括号对都换过后如果长度无变化，说明有不成对的
// 这种方法要进行字符串的查找替换，要生成新的字符串，效率不高
