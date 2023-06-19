package lc0022

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(tmp []byte, l, r int)
	dfs = func(tmp []byte, l, r int) {
		if r == n {
			ans = append(ans, string(tmp))
			return
		}
		if l < n {
			dfs(append(tmp, '('), l+1, r)
		}
		if l > r {
			dfs(append(tmp, ')'), l, r+1)
		}
	}
	dfs(make([]byte, 0, n*2), 0, 0)
	return ans
}
