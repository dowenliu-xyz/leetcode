package main

func longestCommonPrefix(strs []string) string {
	// 纵向扫描
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	l := 0
	for {
		if l >= len(strs[0]) {
			break
		}
		b := strs[0][l]
		for _, str := range strs[1:] {
			if l >= len(str) || str[l] != b {
				return strs[0][:l]
			}
		}
		l++
	}
	return strs[0][:l]
}
