package lc0013

func romanToInt(s string) int {
	mp := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		v := mp[s[i]]
		if i < len(s)-1 && v < mp[s[i+1]] {
			ans -= v
		} else {
			ans += v
		}
	}
	return ans
}
