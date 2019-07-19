package tree

import (
	"fmt"
	"testing"
	Stack "github.com/golang-collections/collections/stack"
)

//构造最大树, 时间和空间复杂度都是O[N]
// 数组 maxTree 定义如下：
// 1、 数组必须没有重复元素
// 2、 MaxTree是一颗二叉树，数组上每一个元素对应树的一个节点
// 3、 树及其子树，最大值必须是root节点

//通过下面两点建立树：
// 1、 每个数的父节点 是它左边第一个比它大的数 和它右边第一个比它大的数中，较小的哪一个
// 2、 如果一个数左边
// 							[3,4,5,1,2]
//							  5
//							/   \
//						   4     2
//                        /		/
//						 3     1

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func PrintMap(bigMap map[*TreeNode]*TreeNode)  {
	for k, v := range bigMap{
		var str1, str2 string
		if k != nil {
			str1 = fmt.Sprintf("%d", k.value)
		}
		if v != nil {
			str2 = fmt.Sprintf("%d", v.value)
		}
		fmt.Printf(" %s -> %s \n", str1, str2)
	}
	fmt.Println("Print end; ============")
}

func setBigMap(bigMap map[*TreeNode]*TreeNode, stack *Stack.Stack)  {
	for {
		node := stack.Pop().(*TreeNode)
		if stack.Len() == 0{
			bigMap[node] = nil
			stack.Push(node)
			break
		}else {
			peekNode := stack.Pop().(*TreeNode)
			if peekNode.value < node.value{
				bigMap[peekNode] = node
				stack.Push(node)
			}else {
				stack.Push(peekNode)
				stack.Push(node)
				bigMap[node] = nil
				bigMap[peekNode] = nil
				break
			}
		}
	}
}


func GetMaxTree(arr []int) (head *TreeNode) {
	length := len(arr)
	node := make([]TreeNode, length)
	for index, ele := range arr {
		node[index].value = ele
	}
	lBigMap := map[*TreeNode]*TreeNode{}
	rBigMap := map[*TreeNode]*TreeNode{}
	stack := Stack.New()
	for index := range node{
		currNode := &node[index]
		stack.Push(currNode)
		setBigMap(rBigMap, stack)
	}
	stack = Stack.New()
	for index := len(node)-1; index >= 0; index -- {
		currNode := &node[index]
		stack.Push(currNode)
		setBigMap(lBigMap, stack)
	}
	for index := range node{
		currNode := &node[index]
		rTreeNode := rBigMap[currNode]
		lTreeNode := lBigMap[currNode]
		if rTreeNode == nil && lTreeNode == nil {
			head = currNode
		}else if rTreeNode == nil {
			if lTreeNode.left == nil {
				lTreeNode.left = currNode
			}else {
				lTreeNode.right = currNode
			}
		}else if lTreeNode == nil {
			if rTreeNode.left == nil {
				rTreeNode.left = currNode
			}else {
				rTreeNode.right = currNode
			}
		}else{
			var parentNode *TreeNode
			if lTreeNode.value > rTreeNode.value {
				parentNode= rTreeNode
			}else {
				parentNode = lTreeNode
			}
			if parentNode.left == nil {
				parentNode.left = currNode
			}else {
				parentNode.right = currNode
			}
		}
	}
	return
}

func Test_GetMaxTree(t *testing.T)  {
	head := GetMaxTree([]int{9,0,1})
	preOrderRecur(head)
}
