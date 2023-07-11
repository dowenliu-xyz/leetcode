package lc0093

import "strings"

func restoreIpAddresses(s string) []string {
	n := len(s)
	var ans []string
	var dfs func(segments []string, i int)
	dfs = func(segments []string, i int) {
		if len(segments) == 4 {
			if i == n {
				ans = append(ans, strings.Join(segments, "."))
			}
			return
		}
		remSegs, remDigs := 4-len(segments), len(s[i:])
		if remDigs < remSegs || remDigs > remSegs*3 {
			return
		}
		for j := 1; j <= 3; j++ {
			if i+j > len(s) {
				break
			}
			if !isValidSegment(s[i : i+j]) {
				continue
			}
			dfs(append(segments, s[i:i+j]), i+j)
		}
	}
	dfs(make([]string, 0, 4), 0)
	return ans
}

func isValidSegment(s string) bool {
	if s == "" || len(s) > 3 {
		return false
	}
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	v := 0
	for i := 0; i < len(s); i++ {
		v *= 10
		v += int(s[i] - '0')
	}
	return v < 256
}
