package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			s = s + strconv.Itoa(node.Val)
			res = append(res, s)
		}
		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}
	travel(root, "")
	return res
}

func main() {

}
