package main

import (
	"fmt"
)

/*
	给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。
	示例 1:
	输入: [[1,1],[2,2],[3,3]]
	输出: 3
	解释:
	^
	|
	|        o
	|     o
	|  o  
	+------------->
	0  1  2  3  4
 */

func abs(a int) int {
	if a >= 0 {
		return a
	}else{
		return -a
	}
}

func maxCommonDivisor(x int, y int) (divisor int) {
	divisor = 1
	for {
		divisor = (x % y)
		if divisor > 0 {
			x = y
			y = divisor
		} else {
			return y
		}
	}
}


func KSlope(pointA []int, pointB []int) string {
	/*
		if pointA[0] == pointB[0] {
			return math.Inf(0)
		}
		return float64(pointA[1] - pointB[1]) / float64(pointA[0] - pointB[0])
	*/
	// 因为会存在斜率超出范围的情况下， float64会超出精度
	if pointA[0] == pointB[0] {
		return  "NAN"
	} else if pointA[1] == pointB[1] {
		return "0"
	}else {
		x := pointA[0] - pointB[0]
		y := pointA[1] - pointB[1]
		flag := ""
		if (x > 0 && y >= 0) || (x < 0 && y < 0) {
			flag = ""
		}else {
			flag = "-"
		}
		x = abs(x)
		y = abs(y)
		commonDivisor := maxCommonDivisor(x, y)
		x = x/commonDivisor
		y = y/commonDivisor
		return fmt.Sprintf("%s%d-%d", flag, y, x)
	}
}

func checkIsSameLine(points [][]int, pA int, mMap map[string]int) (num int) {
	same := 1
	tMax := 0
	for pB := pA + 1; pB < len(points); pB++ {
		pointA := points[pA]
		pointB := points[pB]
		if pointA[0] == pointB[0] && pointA[1] == pointB[1] {
			same++
		}else {
			slope := KSlope(pointA, pointB)
			// 检测斜率是否存在map 中
			slopValue, ok := mMap[slope]
			if ok {
				mMap[slope] = slopValue + 1
			}else {
				mMap[slope] = 1
			}
			if tMax < mMap[slope] {
				tMax = mMap[slope]
			}
		}
	}
	return tMax + same
}

func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}
	num := 0
	for i := 0; i < len(points)-2; i++ {
		mMap := map[string]int{}
		tMax := checkIsSameLine(points, i, mMap)
		if tMax > num {
			num = tMax
		}
	}
	return num
}

func main() {
	//[[3,10],[0,2],[0,2],[3,10]]
	m := []int{1,2,3}
	fmt.Println(m[2:3])
	//fmt.Println(maxPoints([][]int{[]int{1, 1}, []int{2,2},[]int{3,3}}))
	//fmt.Println(maxPoints([][]int{[]int{3,10},[]int{0,2},[]int{0,2}, []int{3,10}}))
}