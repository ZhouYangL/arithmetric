package main

import "fmt"

//给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
//
//示例:
//
//输入: [1,2,3,4]
//输出: [24,12,8,6]
//说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题

func productExceptSelf(nums []int) []int {
	num1 := make([]int, len(nums))
	num2 := make([]int, len(nums))
	for i := range nums{
		if i == 0 {
			num1[i] = 1
		}else {
			num1[i] = num1[i-1]*nums[i-1]
		}
	}
	k := 1
	for i := len(nums) - 1; i >= 0; i-- {
		k *= num1[i] * nums[i]
		num2[len(num2)-1-i] = k
	}
	fmt.Println(num2)
	return num2
}

func main() {
	productExceptSelf([]int{1,2,3,4})
}