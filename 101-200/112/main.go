package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var hasSum func(node *TreeNode, sum int) bool
	hasSum = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}
		sum = sum + node.Val
		if node.Left == nil && node.Right == nil {
			return sum == targetSum
		}
		return hasSum(node.Left, sum) || hasSum(node.Right, sum)
	}
	return hasSum(root, 0)
}

func main() {

}
