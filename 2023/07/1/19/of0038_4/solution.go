package of0038

func permutation(s string) []string {
	n := len(s)
	if n == 0 {
		return nil
	}
	chars := []byte(s)
	var ans []string
	var dfs func(x int)
	dfs = func(x int) {
		if x == n {
			ans = append(ans, string(chars))
			return
		}
		used := map[byte]bool{}
		for i := x; i < n; i++ {
			if used[chars[i]] {
				continue
			}
			used[chars[i]] = true
			chars[i], chars[x] = chars[x], chars[i]
			dfs(x + 1)
			chars[i], chars[x] = chars[x], chars[i]
		}
	}
	dfs(0)
	return ans
}
