package lc0001_1

func twoSum(nums []int, target int) []int {
	//return solution1(nums, target)
	//return solution2(nums, target)
	return solution3(nums, target)
}

func solution1(nums []int, target int) []int {
	// 暴力遍历。注意是返回下标
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1+num2 == target && i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}

func solution2(nums []int, target int) []int {
	// 稍微优化的暴力遍历
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func solution3(nums []int, target int) []int {
	// 查表法。
	// 注意，因为可能出现值相同的元素，所以不能使用一个 map 先扫一遍记录索引，有可能冲突
	table := map[int]int{}
	for i, num := range nums {
		if j, ok := table[target-num]; ok {
			return []int{j, i}
		}
		table[num] = i
	}
	return nil
}
