package main

import "fmt"

/*
	给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。
	返回所有这些可能的句子。
	说明：
	分隔时可以重复使用字典中的单词。
	你可以假设字典中没有重复的单词。
	示例 1：
	输入:
	s = "catsanddog"
	wordDict = ["cat", "cats", "and", "sand", "dog"]
	输出:
	[
	  "cats and dog",
	  "cat sand dog"
	]
	示例 2：
	输入:
	s = "pineapplepenapple"
	wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
	输出:
	[
	  "pine apple pen apple",
	  "pineapple pen apple",
	  "pine applepen apple"
	]
 */

func word_break(s string, mMap map[int][]string, start int, mWordDict map[string]bool) []string {
	value, ok := mMap[start]
	if ok {
		return value
	} else {
		res := []string{}
		if start == len(s) {
			res = append(res, "")
		}
		end := start + 1
		for ; end <= len(s); end++ {
			_, ok := mWordDict[s[start:end]]
			if !ok {
				continue
			}
			list := word_break(s, mMap, end, mWordDict)
			for _, ele := range list {
				tmp := s[start: end]
				if ele == "" {
					res = append(res, tmp)
				}else {
					res = append(res, tmp + " " + ele)
				}
			}
		}
		mMap[start] = res
		return res
	}
}

func wordBreak(s string, wordDict []string) []string {
	mMap := map[int][]string{}
	mWordDict := map[string]bool{}
	for _, ele := range wordDict {
		mWordDict[ele] = true
	}
	return word_break(s, mMap, 0, mWordDict)
}

func main() {
	fmt.Print([]string{"1111", "2222"})
	as := wordBreak("abcd", []string{"ab", "cd", "abc", "d"})
	fmt.Println(as)
}
