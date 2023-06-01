package lc0136_1

func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans = ans ^ num
	}
	return ans
}
