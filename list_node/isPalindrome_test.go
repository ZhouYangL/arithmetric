package list_node

import (
	"testing"
	"fmt"
	Stack "github.com/golang-collections/collections/stack"
)

//判断一个链表是否为回文结构
// 1 -> 2 -> 1 返回 true
// 1 -> 2 -> 2 -> 1 返回true
// 15 -> 6 -> 15 返回true
// 1 -> 2 -> 3 返回false


func isPalindrome(head *Node) (bool) {
	curr := head
	stack := Stack.New()
	for curr != nil {
		stack.Push(curr)
		curr = curr.next
	}
	curr = head
	for stack.Len() > 0 {
		if curr.Value != stack.Peek().(*Node).Value {
			return false
		}
		curr = curr.next
		stack.Pop()
	}
	return true
}

func Test_isPalindrome(t *testing.T)  {
	head := NewNode([]int{1})
	fmt.Println(isPalindrome(head))
}