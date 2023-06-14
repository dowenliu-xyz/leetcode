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
	var dfs func(bytes []byte, idx int)
	dfs = func(bytes []byte, idx int) {
		if idx == n {
			ans = append(ans, string(bytes))
			return
		}
		chars := mp[digits[idx]]
		for i := 0; i < len(chars); i++ {
			dfs(append(bytes, chars[i]), idx+1)
		}
	}
	dfs(make([]byte, 0, n), 0)
	return ans
}
