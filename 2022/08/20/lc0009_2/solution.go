package main

import "strconv"

func isPalindrome(x int) bool {
	//return solution1(x)
	//return solution2(x)
	return solution3(x)
}

func solution1(x int) bool {
	// 先转成字符串再判断是否回文串
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func solution2(x int) bool {
	// 计算反转数。可能溢出，但测试用例可以过。
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed, tmp := 0, x
	for tmp > 0 {
		reversed = reversed*10 + tmp%10 // 这里可能溢出
		tmp /= 10
	}
	return reversed == x
}

func solution3(x int) bool {
	// 计算一半反转
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	for reversed < x {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return reversed == x || reversed/10 == x
}
