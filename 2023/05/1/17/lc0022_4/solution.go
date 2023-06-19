package lc0022

func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(ans *[]string, left, right, n int, data []byte)
	dfs = func(ans *[]string, left, right, n int, data []byte) {
		if left == right && left == n {
			*ans = append(*ans, string(data))
			return
		}
		if left < n {
			dfs(ans, left+1, right, n, append(data, '('))
		}
		if left > right {
			dfs(ans, left, right+1, n, append(data, ')'))
		}
	}
	dfs(&ans, 0, 0, n, nil)
	return ans
}
