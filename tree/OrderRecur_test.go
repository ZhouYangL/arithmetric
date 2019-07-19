package tree

import "fmt"

//遍历二叉树

//递归遍历
//前序遍历
func preOrderRecur(head *TreeNode)  {
	if head == nil {
		return
	}
	fmt.Print(head.value, " -> ")
	preOrderRecur(head.left)
	preOrderRecur(head.right)
}

//中序遍历
func InOrderRecur(head *TreeNode)  {
	if head == nil {
		return
	}
	InOrderRecur(head.left)
	fmt.Print(head.value,  " -> ")
	InOrderRecur(head.right)
}

//后序遍历
func PosOrderRecur(head *TreeNode)  {
	if head == nil {
		PosOrderRecur(head.left)
		PosOrderRecur(head.right)
		fmt.Print(head.value, " -> ")
	}
}



//非递归遍历
//前序
//中序
//后序
