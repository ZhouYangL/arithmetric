package main

import "fmt"

/*
	给定一个非负整数数组，你最初位于数组的第一个位置。
	
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	
	你的目标是使用最少的跳跃次数到达数组的最后一个位置。
	
	示例:
	
	输入: [2,3,1,1,4]
	输出: 2
	解释: 跳到最后一个位置的最小跳跃数是 2。
	     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
*/

func max(a int, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func jump(nums []int) int {
	end := 0
	count := 0
	pos := 0
	for i := 0; i < len(nums) -1; i++ {
		pos = max(pos, i + nums[i])
		if end == i {
			end = pos
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
}