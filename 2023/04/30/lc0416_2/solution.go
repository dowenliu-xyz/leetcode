package lc0416_2

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false // 空集和只含一个元素的集体无法分割
	}
	sum, max := 0, 0 // 因为已经保证集合中元素都是正整数，所以 max 初始值可以用 0 否则最好使用 math.MinInt
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 == 1 {
		return false // 所有元素和为奇数，无法平均均分成两个集合
	}
	target := sum / 2
	if max == target {
		return true // 刚好最大元素是所有元素和的一半
	}
	if max > target {
		return false // 最大元素大于所有元素和的一半，其他元素的和必定无法达到元素和的一半
	}
	dp := make([]bool, target+1) // dp[j] 表示能否实现元素和为 j; 反直觉，这里如果使用 map[int]bool，不论是执行时间还是内存消耗都比 []bool 高
	dp[0] = true                 // 元素和为0总能实现：空集
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- { // 这里要逆序执行。如果升序处理，会把后面计算依赖的值覆盖
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}
