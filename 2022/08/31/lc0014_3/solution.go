package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s0 := strs[0]
	for i := range s0 {
		for j := 1; j < len(strs); j++ {
			s := strs[j]
			if i == len(s) || s0[i] != s[i] {
				return s0[:i]
			}
		}
	}
	return s0
}
