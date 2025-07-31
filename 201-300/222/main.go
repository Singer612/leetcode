package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 1
	var count func(node *TreeNode)
	count = func(node *TreeNode) {
		if node.Left != nil {
			count(node.Left)
			sum++
		}
		if node.Right != nil {
			count(node.Right)
			sum++
		}
	}
	count(root)
	return sum
}

func main() {

}
