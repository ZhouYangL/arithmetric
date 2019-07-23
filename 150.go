package main

import (
	"fmt"
	"strconv"
)

/*
	根据逆波兰表示法，求表达式的值。
	有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
	说明：
	整数除法只保留整数部分。
	给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
	示例 1：
		输入: ["2", "1", "+", "3", "*"]
		输出: 9
		解释: ((2 + 1) * 3) = 9
	示例 2：
		输入: ["4", "13", "5", "/", "+"]
		输出: 6
		解释: (4 + (13 / 5)) = 6
	示例 3：
		输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
		输出: 22
		解释:
		  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
		= ((10 * (6 / (12 * -11))) + 17) + 5
		= ((10 * (6 / -132)) + 17) + 5
		= ((10 * 0) + 17) + 5
		= (0 + 17) + 5
		= 17 + 5
		= 22
 */

func calNum(first int, second int, op string) int {
	cal := 0
	switch op {
	case "+":
		cal = first + second
		break
	case "-":
		cal = first - second
		break
	case "*":
		cal = first * second
		break
	case "/":
		cal = first / second
		break
	}
	return cal
}

func evalRPN(tokens []string) int {
	index := 0
	stack := make([]int, len(tokens))
	var nMap150 = map[string]bool{"+": true, "-": true, "*": true, "/": true}
	cal := 0
	for i := range tokens {
		if nMap150[tokens[i]] {
			cal = calNum(stack[index-2], stack[index-1], tokens[i])
			stack[index-2] = cal
			index -= 1
		}else {
			stack[index], _ = strconv.Atoi(tokens[i])
			index++
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4","-2","/","2","-3","-","-"}))
}