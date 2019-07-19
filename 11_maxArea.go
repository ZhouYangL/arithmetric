package main

import "fmt"

//	给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
//  (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//	说明：你不能倾斜容器，且 n 的值至少为 2。
//	示例:
//	输入: [1,8,6,2,5,4,8,3,7]
//	输出: 49
//
//  f(i, j) = max(min(arr[i], arr[j]) * (j-i))
//

func min(a int, b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}

func maxArea(arr []int) int {
	low := 0
	height := len(arr) - 1
	water := 0
	for ; low < height;  {
		tWater := ( height - low ) * min(arr[low], arr[height])
		if tWater > water {
			water = tWater
		}
		if arr[low] < arr[height] {
			low++
		}else {
			height--
		}
	}
	return water
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}