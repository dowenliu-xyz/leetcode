package main

import "strconv"

func isPalindrome(x int) bool {
	//return solution1(x)
	//return solution2(x)
	return solution3(x)
}

func solution1(x int) bool {
	// 先转字符串再判断回文串
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func solution2(x int) bool {
	// 计算回文数
	if x < 0 {
		// 负数一定不是回文数
		return false
	}
	if x%10 == 0 && x != 0 {
		// 以0结尾的回文数只有0
		return false
	}
	var z int
	cur := x
	for cur > 0 {
		z = z*10 + cur%10 // z 有可能溢出。不过提交通过
		cur /= 10
	}
	return x == z
}

func solution3(x int) bool {
	// 计算一半回文数
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return x == revertedNumber || x == revertedNumber/10
}
