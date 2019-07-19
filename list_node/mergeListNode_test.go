package list_node

import "testing"

// 合并连个有序的单链表
// 0->2->3->7->null
// 1-3->5-9->null

func mergeListNode(h1 *Node, h2 *Node) (node *Node) {
	if h1 == nil {
		return h2
	} else if h2 == nil {
		return h1
	}
	data := make([]int, h1.length() + h2.length())
	index := 0
	curl1 := h1
	curl2 := h2
	for curl1 != nil && curl2 != nil {
		if curl1.Value < curl2.Value {
			data[index] = curl1.Value
			curl1 = curl1.next
		} else {
			data[index] = curl2.Value
			curl2 = curl2.next
		}
		index ++
	}
	for curl2 != nil {
		data[index] = curl2.Value
		curl2 = curl2.next
		index ++
	}
	for curl1 != nil {
		data[index] = curl1.Value
		curl1 = curl1.next
		index ++
	}
	node = NewNode(data)
	return
}


func Test_2_mergeListNode(t *testing.T)  {
	h1 := NewNode([]int{0,2,4,7,9})
	h2 := NewNode([]int{1,2,4,6,8})
	node := mergeListNode(h1, h2)
	node.PrintNode()
	
	//------------------------------
	println("\nTest_2_mergeListNode: -------------")
	var h3 *Node
	var h4 *Node
	node = mergeListNode(h3, h4)
	node.PrintNode()
	
	println("\nTest_3_mergeListNode: -------------")
	var h5 *Node
	h6 := NewNode([]int{1,3,5,6,8})
	node = mergeListNode(h5, h6)
	node.PrintNode()
}
