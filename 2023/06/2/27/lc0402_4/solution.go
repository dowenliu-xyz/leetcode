package lc0402

func removeKdigits(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}
	ans := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		for k > 0 && len(ans) > 0 && ans[len(ans)-1] > num[i] {
			ans = ans[:len(ans)-1]
			k--
		}
		ans = append(ans, num[i])
	}
	if k > 0 {
		ans = ans[:len(ans)-k]
	}
	for len(ans) > 0 && ans[0] == '0' {
		ans = ans[1:]
	}
	if len(ans) == 0 {
		return "0"
	}
	return string(ans)
}
