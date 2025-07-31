package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var getHeight func(node *TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := getHeight(node.Left)
		rightHeight := getHeight(node.Right)
		if leftHeight == -1 || rightHeight == -1 || math.Abs(float64(leftHeight-rightHeight)) > 1 {
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}
	return !(getHeight(root) == -1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
