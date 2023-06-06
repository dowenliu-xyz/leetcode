package lc0014_2

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	for i := 0; i < len(first); i++ {
		for j := 1; j < len(strs); j++ {
			s := strs[j]
			if i == len(s) || first[i] != s[i] {
				return s[:i]
			}
		}
	}
	return first
}
