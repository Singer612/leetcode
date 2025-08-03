package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	// 保留前一个节点的指针
	var pre *TreeNode
	minVal := math.MaxInt64
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if pre != nil && node.Val-pre.Val < minVal {
			minVal = node.Val - pre.Val
		}
		pre = node
		travel(node.Right)
	}
	travel(root)
	return minVal
}

func main() {

}
