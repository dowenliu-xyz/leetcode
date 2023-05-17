package lc0217_4

func containsDuplicate(nums []int) bool {
	seen := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}
