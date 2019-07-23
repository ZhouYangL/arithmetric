package main

import (
	"fmt"
)

/*
	给定一个字符串，逐个翻转字符串中的每个单词。
	示例 1：
		输入: "the sky is blue"
		输出: "blue is sky the"
	示例 2：
		输入: "  hello world!  "
		输出: "world! hello"
		解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
	示例 3：
		输入: "a good   example"
		输出: "example good a"
		解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
	说明：
		无空格字符构成一个单词。
		输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
		如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 */

func addString(result string, add string) string {
	if result == "" {
		result = add
	}else if add != "" {
		result = result + " " + add
	}
	return result
}
func reverseWords(s string) string {
	result := ""
	end := len(s)
	start := len(s)-1
	for ; start >= 0; start-- {
		if s[start] == ' ' {
			if s[start+1:end] == ""{
				end = start
				continue
			}else {
				result= addString(result, s[start+1:end])
				end = start
			}
		}
	}
	if start < end{
		result= addString(result, s[start+1:end])
	}
	return result
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))
}


