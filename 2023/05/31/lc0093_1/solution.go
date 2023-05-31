package lc0093_1

import "strings"

func restoreIpAddresses(s string) []string {
	var ans []string
	search([]byte(s), &ans, make([][]byte, 0, 4), 0)
	return ans
}

func search(s []byte, ans *[]string, parts [][]byte, i int) {
	if len(parts) == 4 {
		*ans = append(*ans, join(parts))
		return
	}
	if len(s[i:]) < 4-len(parts) || len(s[i:]) > 3*(4-len(parts)) {
		return
	}
	if len(parts) == 3 {
		if isValidPart(s[i:]) {
			search(s, ans, append(parts, s[i:]), 0)
		}
		return
	}
	for j := 1; j <= 3; j++ {
		if i+j > len(s) || !isValidPart(s[i:i+j]) {
			continue
		}
		search(s, ans, append(parts, s[i:i+j]), i+j)
	}
}

func isValidPart(s []byte) bool {
	n := len(s)
	if n > 1 && s[0] == '0' { // 除了 0 之外有前导 0，无效
		return false
	}
	tmp := 0
	for i := 0; i < n; i++ {
		tmp = tmp*10 + int(s[i]-'0')
	}
	return tmp <= 255
}

func join(parts [][]byte) string {
	sb := strings.Builder{}
	for i, part := range parts {
		if i > 0 {
			sb.WriteByte('.')
		}
		sb.Write(part)
	}
	return sb.String()
}
