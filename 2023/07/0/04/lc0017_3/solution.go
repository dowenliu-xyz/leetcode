package lc0017

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}
	mp := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi",
		'5': "jkl", '6': "mno", '7': "pqrs",
		'8': "tuv", '9': "wxyz",
	}
	var ans []string
	var dfs func(tmp []byte)
	dfs = func(tmp []byte) {
		if len(tmp) == n {
			ans = append(ans, string(tmp))
			return
		}
		str := mp[digits[len(tmp)]]
		for i := 0; i < len(str); i++ {
			dfs(append(tmp, str[i]))
		}
	}
	dfs(make([]byte, 0, n))
	return ans
}
