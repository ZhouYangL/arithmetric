package main

import "fmt"

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//示例 1:
//
//输入: 121
//输出: true


func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	sum := 0
	curl := x
	for curl != 0 {
		pop := curl % 10
		sum = sum * 10 + pop
		curl /= 10
	}
	if sum == x {
		return true
	}else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(-121))
}