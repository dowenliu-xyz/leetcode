package lc0093

import "strings"

func restoreIpAddresses(s string) []string {
	var ans []string
	search([]byte(s), &ans, make([][]byte, 0, 4), 0)
	return ans
}

func search(s []byte, ans *[]string, parts [][]byte, i int) {
	if len(parts) == 3 {
		if isValidPart(s[i:]) {
			*ans = append(*ans, join(append(parts, s[i:])))
		}
		return
	}
	remainLen := len(s[i:])
	remainParts := 4 - len(parts)
	if remainLen < remainParts || remainLen > 3*remainParts {
		return
	}
	for j := 1; j <= 3; j++ {
		k := i + j
		if k >= len(s) {
			break
		}
		if isValidPart(s[i:k]) {
			search(s, ans, append(parts, s[i:k]), k)
		}
	}
}

func isValidPart(s []byte) bool {
	n := len(s)
	if n == 0 || n > 3 {
		return false
	}
	if n > 1 && s[0] == '0' {
		return false
	}
	v := 0
	for i := 0; i < n; i++ {
		v = v*10 + int(s[i]-'0')
	}
	return v < 256
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
