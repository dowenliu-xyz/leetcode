package lc0001_5

func twoSum(nums []int, target int) []int {
	mem := make(map[int]int, len(nums))
	for i := range nums {
		j, ok := mem[target-nums[i]]
		if ok {
			return []int{j, i}
		}
		mem[nums[i]] = i
	}
	return nil
}
