package list_node

import (
	"errors"
	"fmt"
	"testing"
)

//链表 1->2->3->null 给定要删除节点2， 在不知道头节点的情况下删除节点2

type Node struct {
	Value int;
	next *Node;
}


func NewNode(data []int) (head *Node) {
	head = nil
	var end *Node
	for _, value := range data{
		if head == nil{
			head = &Node{
				Value:value,
				next:nil,
			}
			end = head
		} else {
			tmp := &Node{
				Value: value,
				next:nil,
			}
			end.next = tmp
			end = tmp
		}
	}
	return
}


func (node *Node)PrintNode()  {
	tmpNode := node
	all := ""
	for tmpNode != nil{
		all += fmt.Sprintf("%d -> ", tmpNode.Value)
		tmpNode = tmpNode.next
	}
	fmt.Println(all)
}


func (node *Node)length() (count int) {
	tmp := node
	for tmp != nil {
		count ++
		tmp = tmp.next
	}
	return
}


func removeNodeWired(node *Node) (error) {
	next := node.next
	if next == nil {
		return errors.New("don't remove last node")
	}
	node.Value = next.Value
	node.next = next.next
	return nil
}

func Test_removeNodeWired(t *testing.T)  {
	head := NewNode([]int{1,2,3,4,5,6,7})
	Node3 := head.next.next
	removeNodeWired(Node3)
	head.PrintNode()
}
