package list_node

import (
	"fmt"
	"sort"
	"testing"
)

//向有序的环状链表随机插入值

type CircleListNode struct {
	value int
	next *CircleListNode
}

func NewCircleNode(data []int) (head *CircleListNode) {
	head = nil
	var end *CircleListNode
	for _, value := range data{
		if head == nil{
			head = &CircleListNode{value:value, next:nil}
			end = head
			end.next = head
		} else {
			tmp := &CircleListNode{value: value, next:nil}
			end.next = tmp
			end = tmp
			end.next = head
		}
	}
	return
}


func (node *CircleListNode)PrintNode(isSort bool)  {
	if node == nil {
		return
	}
	data := make([]int, node.length())
	index := 0
	head := node
	tmpNode := node.next
	data[index] = head.value
	for  head != tmpNode{
		index++
		data[index] = tmpNode.value
		tmpNode = tmpNode.next
	}
	if isSort{
		sort.Ints(data)
	}
	all := ""
	for _, value := range data{
		all += fmt.Sprintf("%d -> ", value)
	}
	fmt.Println(all)
}


func (node *CircleListNode)length() (count int) {
	if node == nil {
		return
	}
	tmp := node.next
	head := node
	count ++
	for head != tmp {
		count ++
		tmp = tmp.next
	}
	return
}

func (head *CircleListNode)insertNode(num int) (*CircleListNode){
	node :=&CircleListNode{value: num, next:nil}
	if head == nil {
		head = node
		head.next = head
	}else {
		node.next = head.next
		head.next = node
	}
	return head
}
func insertNum(head *CircleListNode, num int) (newHead *CircleListNode) {
	if head == nil {
		newHead = head.insertNode(num)
		return
	}
	newHead = head
	pre := head.next
	curl := head
	if pre == curl {
		head.insertNode(num)
		return
	}

	min := head
	for flag:=true; flag || head != curl; flag=false {
		if (curl.value + pre.value)/2 - num >= 0{
			curl.insertNode(num)
			return
		}
		curl = pre
		pre = pre.next
		if min.value > curl.value{
			min = curl
		}
	}
	if curl == head {
		min.insertNode(num)
	}
	return
}

func Test_insertNum(t *testing.T)  {
	fmt.Println("\n测试1---------------")
	head :=NewCircleNode([]int{4,3,2,1})
	head.PrintNode(false)
	
	fmt.Println("\n测试2---------------")
	insertNum(head, 6)
	insertNum(head, 5)
	head.PrintNode(false)
	
	fmt.Println("\n测试3---------------")
	head =NewCircleNode([]int{4})
	insertNum(head, 6)
	head.PrintNode(false)
	
	fmt.Println("\n测试4---------------")
	head =NewCircleNode([]int{})
	head = head.insertNode(5)
	head.PrintNode(false)
	insertNum(head, 6)
	head.PrintNode(false)
}