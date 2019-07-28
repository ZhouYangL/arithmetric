package main

/*
	给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
	说明: 叶子节点是指没有子节点的节点。
	示例: 
	给定如下二叉树，以及目标和 sum = 22
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
	返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
 */

func OrderPath(root *TreeNode, sum int) (int, bool) {
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return sum, true
	}
	if root.Left != nil{
		if _, left := OrderPath(root.Left, sum); left {
			return 0, true
		}
	}
	if root.Right != nil{
		if _, right := OrderPath(root.Right, sum); right {
			return 0, true
		}
	}
	return sum, false
}
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	_, result := OrderPath(root, sum)
	return result
}
