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
		used := map[byte]struct{}{}
		for i := x; i < n; i++ {
			if _, ok := used[chars[i]]; ok {
				continue
			}
			used[chars[i]] = struct{}{}
			chars[i], chars[x] = chars[x], chars[i]
			dfs(x + 1)
			chars[i], chars[x] = chars[x], chars[i]
		}
	}
	dfs(0)
	return ans
}
