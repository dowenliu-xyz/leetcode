package main

func longestCommonPrefix(strs []string) string {
	// 纵向扫描
	if len(strs) == 0 {
		return ""
	}
	var first string
	first, strs = strs[0], strs[1:]
	for i := 0; i < len(first); i++ {
		for _, str := range strs {
			if i >= len(str) || first[i] != str[i] {
				return first[:i]
			}
		}
	}
	return first
}
