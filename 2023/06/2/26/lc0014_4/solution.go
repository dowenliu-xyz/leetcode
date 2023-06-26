package lc0014

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 1 {
		return strs[0]
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < n; j++ {
			if i == len(strs[j]) || c != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
