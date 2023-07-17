package lc0043

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ans := make([]byte, m+n)
	for i := m - 1; i >= 0; i-- {
		x := num1[i] - '0'
		for j := n - 1; j >= 0; j-- {
			y := num2[j] - '0'
			k := i + j + 1
			tmp := x * y
			for tmp != 0 {
				ans[k] += tmp
				tmp = ans[k] / 10
				ans[k] %= 10
				k--
			}
		}
	}
	for len(ans) > 0 && ans[0] == 0 {
		ans = ans[1:]
	}
	for i := range ans {
		ans[i] += '0'
	}
	return string(ans)
}
