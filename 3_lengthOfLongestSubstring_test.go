package main

import (
	"testing"
	"fmt"
)

// Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	bigMap := map[rune]int{}
	max := 0
	start := 0
	end := 0
	for index, ch := range s {
		result, ok := bigMap[ch]
		end = index
		if ok && start <= result {
			tmax := end - start
			start = result + 1
			if tmax > max {
				max = tmax
			}
		}
		bigMap[ch] = index
	}
	if max < end - start + 1{
		max = end - start + 1
	}
	return max
}

func Test_lengthOfLongestSubstring(t *testing.T)  {
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
