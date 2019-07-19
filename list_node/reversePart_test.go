package list_node

import (
	"testing"
)

//反转部分单链表
func findNode(head *Node, to int) (h *Node, fpos *Node) {
	h = head
	for h != nil{
		if h.Value == to {
			break
		}
		fpos = h
		h = h.next
	}
	return
}

func ReversePart(head *Node, from int, to int) (*Node){
	h1, fpos := findNode(head, from)
	h2, epos := findNode(head, to)

	if h1 == nil || h2 == nil || h1 == h2 {
		return head
	}
	epos = h2.next
	if h1 == head {
		fpos = &Node{next:head}
		head = h2
	}
	ptpos := h1
	itpos := h1.next
	for ptpos != h2 && itpos != nil {
		otpos := itpos.next
		itpos.next = ptpos
		ptpos = itpos
		itpos = otpos
	}
	fpos.next = h2
	h1.next = epos
	return head
}

func Test_Reverse(t *testing.T)  {
	head := NewNode([]int{1,2,3,4,5,6,7})
	head.PrintNode()
	head = ReversePart(head, 6, 8)
	head.PrintNode()
}