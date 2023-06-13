package of0046

import (
	"strconv"
)

func translateNum(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	//dp := make([]int, n+1)
	//dp[0], dp[1] = 1, 1
	p, q := 1, 1
	for i := 2; i <= n; i++ {
		two := str[i-2 : i]
		if "10" <= two && two <= "25" {
			//dp[i] = dp[i-2] + dp[i-1]
			p, q = q, p+q
		} else {
			//dp[i] = dp[i-1]
			p = q
		}
	}
	//return dp[n]
	return q
}
