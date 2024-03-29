package main

import "fmt"

/*
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
	
	上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。
	示例:
	输入: [0,1,0,2,1,0,1,3,2,1,2,1]
	输出: 6
 */

func max46(a int, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}


func trap(height []int) int {
	maxHeight := 0
	maxIndex := 0
	for index, ele := range height {
		if maxHeight < ele {
			maxHeight = ele
			maxIndex = index
		}
	}
	tMaxHeight := 0
	sum := 0
	for i := 0; i < maxIndex; i++ {
		ele := height[i]
		tMaxHeight = max46(ele, tMaxHeight)
		if tMaxHeight > ele {
			sum = sum + (tMaxHeight - ele)
		}
	}
	tMaxHeight = 0
	for i := len(height) - 1; i > maxIndex; i-- {
		ele := height[i]
		tMaxHeight = max46(ele, tMaxHeight)
		if tMaxHeight > ele {
			sum = sum + (tMaxHeight-ele)
		}
	}
	return sum
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
