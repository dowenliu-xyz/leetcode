package lc0560

func subarraySum(nums []int, k int) int {
	// pre: 当前已累积前缀和。mp：当前已知前缀和出现次数
	ans, pre, mp := 0, 0, make(map[int]int)
	for _, num := range nums {
		mp[pre]++
		pre += num
		if count, ok := mp[pre-k]; ok {
			ans += count
		}
	}
	return ans
}
