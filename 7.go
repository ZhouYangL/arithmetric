package main

import (
	"fmt"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//
//输入: 123
//输出: 321
// 示例 2:
//
//输入: -123
//输出: -321
//示例 3:
//
//输入: 120
//输出: 21

const maxNum = 1 << 31 - 1
const minNum = -1 << 31
func reverse(x int) int {
	s := 0
	for x != 0 {
		m := x % 10
		x = x / 10
		if s > maxNum / 10 || (s == maxNum / 10 && m > 7){
			return 0
		}
		if s < minNum / 10 || (s == maxNum / 10 && m < -8){
			return 0
		}
		s = s * 10 + m
	}
	return s
}

func main()  {
	fmt.Println(minNum)
	fmt.Println(reverse(123))
}
