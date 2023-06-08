package lc0038

func permutation(s string) []string {
	n := len(s)
	if n == 0 {
		return nil
	}
	bs := []byte(s)
	var ans []string
	var search func(x int)
	search = func(x int) {
		if x == n-1 {
			ans = append(ans, string(bs))
			return
		}
		mem := make(map[byte]struct{})
		for i := x; i < n; i++ {
			if _, ok := mem[bs[i]]; ok {
				continue
			}
			mem[bs[i]] = struct{}{}
			bs[i], bs[x] = bs[x], bs[i]
			search(x + 1)
			bs[i], bs[x] = bs[x], bs[i]
		}
	}
	search(0)
	return ans
}
