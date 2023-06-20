package lc0560

func subarraySum(nums []int, k int) int {
	pre, ans, mp := 0, 0, make(map[int]int)
	for _, num := range nums {
		mp[pre]++
		pre += num
		if count, ok := mp[pre-k]; ok {
			ans += count
		}
	}
	return ans
}
