package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes1(root.Left) + countNodes1(root.Right) + 1
}

// 利用完全二叉树的特性，如果是满二叉树直接返回2的depth次方-1，可以这样做的原因是总可以找到一颗满二叉树
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 判断左右两侧子树深度是否相等，如果相等可以直接节点数量
	leftDepth, rightDepth := 0, 0
	left := root.Left
	right := root.Right
	// 一直找到左子树末尾
	for left != nil {
		leftDepth++
		left = left.Left
	}
	// 右子树深度
	for right != nil {
		rightDepth++
		right = right.Right
	}
	if root.Left == root.Right {
		return 2<<leftDepth - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func main() {

}
