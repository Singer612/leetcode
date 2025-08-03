package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var isValid func(node *TreeNode, lower, upper int) bool
	isValid = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		return isValid(node.Left, lower, node.Val) && isValid(node.Right, node.Val, upper)
	}
	return isValid(root, math.MinInt, math.MaxInt)
}

func main() {

}
