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
		count2[s2[i]-'a']++
	}
	for i := m; i < n; i++ {
		if count1 == count2 {
			return true
		}
		count2[s2[i]-'a']++
		count2[s2[i-m]-'a']--
	}
	return count1 == count2
}
