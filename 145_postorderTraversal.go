package main

/*
		输入: [1,null,2,3]
	   1
		\
		 2
		/
	   3
	
	输出: [3,2,1]
	进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func PosOrderRecur(root *TreeNode, recode []int)  []int {
	if root == nil {
		return recode
	}else {
		recode = PosOrderRecur(root.Left, recode)
		recode = PosOrderRecur(root.Right, recode)
		recode = append(recode, root.Val)
		return recode
	}
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	return PosOrderRecur(root, result)
}
