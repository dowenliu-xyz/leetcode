package lc0567

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	var (
		count1 [26]int
		count2 [26]int
	)
	for i := 0; i < m; i++ {
		count1[s1[i]-'a']++
	}
	for r := 0; r < n; r++ {
		count2[s2[r]-'a']++
		if r < m-1 {
			continue
		}
		if r >= m {
			count2[s2[r-m]-'a']--
		}
		if count1 == count2 {
			return true
		}
	}
	return false
}
