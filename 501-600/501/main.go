package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var pre *TreeNode
	res := make([]int, 0)
	count, maxCount := 1, math.MinInt
	var find func(node *TreeNode)
	find = func(node *TreeNode) {
		if node == nil {
			return
		}
		find(node.Left)
		if pre != nil && pre.Val == node.Val {
			count++
		} else if pre != nil && pre.Val != node.Val {
			count = 1
		} else {
			res = append(res, node.Val)
		}
		if count > maxCount {
			res = nil
			maxCount = count
			res = append(res, node.Val)
		} else if count == maxCount {
			res = append(res, node.Val)
		}
		pre = node
		find(node.Right)
	}
	find(root)
	return res
}

func main() {

}
