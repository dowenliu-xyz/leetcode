package lc0013_3

func romanToInt(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dict := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	ans := 0
	for i := 0; i < n; i++ {
		if i < n-1 && dict[s[i]] < dict[s[i+1]] {
			ans -= dict[s[i]]
			continue
		}
		ans += dict[s[i]]
	}
	return ans
}
