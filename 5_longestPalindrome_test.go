package main

import "math"

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"
//

func exportCenter(s string, left int , right int)  int {
	for ;left >= 0 && right < len(s) && left <= right && s[left] == s[right]; {
		left--
		right++
	}
	return right - left - 1
}

func longestPalindrome(s string) string {
	if s == ""{
		return ""
	}
	end, start := 0, 0
	for i := range s {
		len1 := exportCenter(s, i, i)
		len2 := exportCenter(s, i, i+1)
		len3 := math.Max(float64(len1), float64(len2))
		if len3 > float64(end - start){
			start = i - int(len3 -1 ) / 2
			end = i + int(len3) / 2
		}
	}
	return s[start: end+1]
}



