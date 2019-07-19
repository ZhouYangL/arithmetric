package main

// 给定 [1,2,4,7], 给定数字 8， 输出[0, 3]
import (
	"testing"
	"fmt"
)

func twoSum(nums []int, target int) []int {
	bigMap := map[int]int{}
	for index, ele := range nums {
		other := target - ele
		result, ok := bigMap[other]
		if ok {
			return []int{result, index}
		}
		bigMap[ele] = index
	}
	return []int{0,0}
}

func Test_twoSum(t *testing.T)  {
	fmt.Println(twoSum([]int{3,2,4}, 6))
}