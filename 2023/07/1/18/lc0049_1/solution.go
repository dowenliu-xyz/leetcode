package lc0049

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		var key [26]int
		for i := 0; i < len(str); i++ {
			key[str[i]-'a']++
		}
		mp[key] = append(mp[key], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, group := range mp {
		ans = append(ans, group)
	}
	return ans
}
