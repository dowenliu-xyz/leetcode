package of0046

import "strconv"

func translateNum(num int) int {
	str := strconv.Itoa(num) // 转为字符串
	n := len(str)            // 字符串长度一定不为 0 ，数字转字符串的特性保证了这一点
	// 在翻译的过程中，对到达的每一位，可能是从前一位之后，翻译当前一位到达，也可能是前前两位，一次翻译两位到达（如果这两位可以被合法的翻译
	//*/
	dp := make([]int, n+1) // dp 数组，dp[i] 表示翻译到第 i 位（对应索引下标 i - 1）数字时可能的方法数
	dp[0], dp[1] = 1, 1    // 特殊值，翻译 0 位或第 1 位数字时，都只有一种可能的方法
	/*/
	p, q := 1, 1
	//*/
	for i := 2; i <= n; i++ { // 从第 2 位数字开始考察
		two := str[i-2 : i] // 第 i-1 和第 i 位数字组成的子串，即 str[i-2] 和 str[i-1] 组成的子串
		if "10" <= two && two <= "25" {
			// 这两位数组子串有两种翻译方法
			//*/
			dp[i] = dp[i-2] + dp[i-1]
			/*/
			p, q = q, p + q
			//*/
		} else {
			// 只能是第 i 位数字自己一位进行翻译
			//*/
			dp[i] = dp[i-1]
			/*/
			p = q
			//*/
		}
	}
	//*/
	return dp[n] // 最终结果就在 dp[n] 中
	/*/
	return q
	//*/
}
