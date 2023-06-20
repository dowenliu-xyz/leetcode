package of0005

func replaceSpace(s string) string {
	ans := make([]byte, 0, len(s)*3)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans = append(ans, "%20"...)
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
