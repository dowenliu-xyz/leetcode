package lc0049

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for i := range strs {
		var key [26]int
		for j := 0; j < len(strs[i]); j++ {
			key[strs[i][j]-'a']++
		}
		mp[key] = append(mp[key], strs[i])
	}
	ans := make([][]string, 0, len(mp))
	for _, group := range mp {
		ans = append(ans, group)
	}
	return ans
}
