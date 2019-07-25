package main

import "fmt"

/*
	使用队列实现栈的下列操作：
		push(x) -- 元素 x 入栈
		pop() -- 移除栈顶元素
		top() -- 获取栈顶元素
		empty() -- 返回栈是否为空
	注意:
		你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
		你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
		你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

 */


type MyStack struct {
	Stack []int
	Index int
}


/** Initialize your data structure here. */
func Constructor11() MyStack {
	return MyStack{
		Stack: make([]int, 0),
		Index: 0,
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	if this.Index < cap(this.Stack){
		this.Stack[this.Index] = x
	}else {
		this.Stack = append(this.Stack, x)
	}
	this.Index++
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ele := this.Stack[this.Index-1]
	this.Index--
	return ele
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Stack[this.Index-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Index <= 0
}

func main() {
	//["MyStack","push","pop","push","top"]
	//[[],[1],[],[2],[]]
	a :=  Constructor()
	fmt.Println(a)
	a.Push(1)
	fmt.Println(a)
	a.Pop()
	fmt.Println(a)
	a.Push(2)
	fmt.Println(a)
	fmt.Println(a.Top())
}