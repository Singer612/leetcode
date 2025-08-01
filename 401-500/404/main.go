package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var sum1 func(root *TreeNode, flag bool) int
	sum1 = func(root *TreeNode, flag bool) int {
		// 1.终止条件
		if root.Left == nil && root.Right == nil && flag == true {
			return root.Val
		}
		// 2.每层递归
		leftSum, rightSum := 0, 0
		if root.Left != nil {
			leftSum = sum1(root.Left, true)
		}
		if root.Right != nil {
			flag = false
			rightSum = sum1(root.Right, false)
		}
		// 3.返回信息
		return leftSum + rightSum
	}
	return sum1(root, false)
}

func buildTree() *TreeNode {
	// 创建节点
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func main() {
	tree := buildTree()
	fmt.Println(sumOfLeftLeaves(tree))
}
