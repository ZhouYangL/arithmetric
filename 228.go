package main

import "fmt"

/*
	给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。
	示例 1:
		输入: [0,1,2,4,5,7]
		输出: ["0->2","4->5","7"]
		解释: 0,1,2 可组成一个连续的区间; 4,5 可组成一个连续的区间。
	示例 2:
		输入: [0,2,3,4,6,8,9]
		输出: ["0","2->4","6","8->9"]
		解释: 2,3,4 可组成一个连续的区间; 8,9 可组成一个连续的区间

 */
func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	start := 0
	end := 0
	for end = range nums {
		if nums[end] - nums[start] == end - start {
			continue
		}else {
			if end -1 == start {
				result = append(result, fmt.Sprintf("%d", nums[start]))
			}else {
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
			}
			start = end
		}
	}
	if end  == start {
		result = append(result, fmt.Sprintf("%d", nums[start]))
	}else {
		result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end]))
	}
	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0,1,2,4,5,7}))
	fmt.Println(summaryRanges([]int{0,1,2,4,5,7,8,9}))
}
