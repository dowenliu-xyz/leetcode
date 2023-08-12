package lc0013

func romanToInt(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dict := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	ans := dict[s[n-1]]
	for i := n - 2; i >= 0; i-- {
		v, p := dict[s[i]], dict[s[i+1]]
		if v < p {
			ans -= v
		} else {
			ans += v
		}
	}
	return ans
}
