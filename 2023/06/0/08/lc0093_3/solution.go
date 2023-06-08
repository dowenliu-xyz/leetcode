package lc0093

import "strings"

func restoreIpAddresses(s string) []string {
	var ans []string
	var search func(parts []string, i int)
	search = func(parts []string, i int) {
		if len(parts) == 3 {
			if isValid(s[i:]) {
				ans = append(ans, strings.Join(append(parts, s[i:]), "."))
			}
		}
		rem, remLen := 4-len(parts), len(s[i:])
		if remLen < rem || remLen > rem*3 {
			return
		}
		for j := 1; j <= 3; j++ {
			if i+j >= len(s) {
				break
			}
			if !isValid(s[i : i+j]) {
				continue
			}
			search(append(parts, s[i:i+j]), i+j)
		}
	}
	search(make([]string, 0, 4), 0)
	return ans
}

func isValid(part string) bool {
	n := len(part)
	if n == 0 || n > 3 {
		return false
	}
	if n > 1 && part[0] == '0' {
		return false
	}
	num := 0
	for i := 0; i < n; i++ {
		num *= 10
		num += int(part[i] - '0')
	}
	return num < 256
}
