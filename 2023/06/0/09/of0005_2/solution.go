package of0005

func replaceSpace(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	ans := make([]byte, 0, n*3)
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			ans = append(ans, "%20"...)
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
