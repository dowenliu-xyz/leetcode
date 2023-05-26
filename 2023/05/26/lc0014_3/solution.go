package lc0014_3

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	for i := range strs[0] {
		for j := 1; j < n; j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[j][:i]
			}
		}
	}
	return strs[0]
}
