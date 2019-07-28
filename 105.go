package main

import "fmt"

/*
	根据一棵树的前序遍历与中序遍历构造二叉树。
	注意:
		你可以假设树中没有重复的元素。
	例如，给出
		前序遍历 preorder = [3,9,20,15,7]
		中序遍历 inorder = [9,3,15,20,7]
		返回如下的二叉树：
			3
		   / \
		  9  20
			/  \
		   15   7
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func build(node *TreeNode, index int, preorder []int, inorder []int, mMap map[int]int,left int, right int) (int, *TreeNode) {
	if right < 0 || left >= len(inorder) || left > right{
		return index-1, node
	}
	value := preorder[index]
	mIndex := mMap[value]
	if node == nil {
		node = &TreeNode{Val:value}
		index, _ = build(node, index+1, preorder, inorder, mMap, left, mIndex-1)
		index, _ = build(node, index+1, preorder, inorder, mMap, mIndex+1, right)
		return index, node
	} else if node.Left == nil && mMap[node.Val] > mIndex{
		node.Left = &TreeNode{Val:value}
		index, _ = build(node.Left, index+1, preorder, inorder, mMap, left, mIndex-1)
		index, _ = build(node.Left, index+1, preorder, inorder, mMap, mIndex+1, right)
	} else {
		node.Right = &TreeNode{Val:value}
		index, _ = build(node.Right, index+1, preorder, inorder, mMap, left, mIndex-1)
		index, _ = build(node.Right, index+1, preorder, inorder, mMap, mIndex+1, right)
	}
	return index, node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	mMap := make(map[int]int, len(inorder))
	for index, ele := range inorder {
		mMap[ele] = index
	}
	var root *TreeNode
	_, root = build(root, 0, preorder, inorder, mMap, 0, len(inorder)-1)
	return root
}

func main() {
	fmt.Println(buildTree([]int{1,2}, []int{1,2}))
	fmt.Println(buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
}