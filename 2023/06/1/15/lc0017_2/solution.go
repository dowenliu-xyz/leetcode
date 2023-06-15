package lc0017

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}
	mp := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var ans []string
	var dfs func(tmp []byte, i int)
	dfs = func(tmp []byte, i int) {
		if i == n {
			ans = append(ans, string(tmp))
			return
		}
		chars := mp[digits[i]]
		for j := 0; j < len(chars); j++ {
			dfs(append(tmp, chars[j]), i+1)
		}
	}
	dfs(make([]byte, 0, n), 0)
	return ans
}
