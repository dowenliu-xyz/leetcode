package of0003

func findRepeatNumber(nums []int) int {
	//return solution1(nums)
	return solution2(nums)
}

func solution1(nums []int) int {
	mem := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := mem[num]; ok {
			return num
		}
		mem[num] = struct{}{}
	}
	return -1
}

func solution2(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}
