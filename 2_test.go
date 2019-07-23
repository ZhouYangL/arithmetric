package main

import (
	"testing"
	"fmt"
)

//给定 两个链表
// 2 -> 4 -> 3
// 5 -> 6 -> 4
// output 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ins := 0
	curr1, curr2 := l1, l2;
	var carry int
	var head *ListNode
	for curr1 != nil{
		if curr2 != nil {
			carry = curr2.Val + curr1.Val
			curr2 = curr2.Next
		}else {
			carry =  + curr1.Val
		}
		curr1 = curr1.Next
		if head != nil {
			head = head.Next
			head.Val = (carry + ins) % 10
		}else {
			head = l1
			head.Val = (carry + ins) % 10
		}
		ins = (carry + ins) / 10
	}

	curr1 = curr2
	head.Next = curr2
	for curr1 != nil {
		head = head.Next
		carry := curr1.Val + ins
		head.Val = carry % 10
		ins = carry / 10
		curr1 = curr1.Next
	}
	if ins > 0 {
		head.Next = &ListNode{1, nil}
	}
	return l1
}


func NewNode(data []int) (head *ListNode) {
	head = nil
	var end *ListNode
	for _, value := range data{
		if head == nil{
			head = &ListNode{value, nil}
			end = head
		} else {
			tmp := &ListNode{value, nil}
			end.Next = tmp
			end = tmp
		}
	}
	return
}


func (node *ListNode)PrintNode()  {
	tmpNode := node
	all := ""
	for tmpNode != nil{
		all += fmt.Sprintf("%d -> ", tmpNode.Val)
		tmpNode = tmpNode.Next
	}
	fmt.Println(all)
}


func Test_2(t *testing.T)  {
	l1 := NewNode([]int{9,8})
	l2 := NewNode([]int{1})
	l1.PrintNode()
	l2.PrintNode()
	l3 := addTwoNumbers(l1, l2)
	l3.PrintNode()
	fmt.Println("=================")
	l1 = NewNode([]int{2,4,6})
	l2 = NewNode([]int{5,6,4})
	l1.PrintNode()
	l2.PrintNode()
	l3 = addTwoNumbers(l1, l2)
	l3.PrintNode()
	
	fmt.Println("=================")
	l1 = NewNode([]int{2,4,5, 1})
	l2 = NewNode([]int{5,6,4})
	l1.PrintNode()
	l2.PrintNode()
	l3 = addTwoNumbers(l1, l2)
	l3.PrintNode()
	
}