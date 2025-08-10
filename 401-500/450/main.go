package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left != nil && root.Right != nil {

			root.Right.Left = root.Left
			return root.Right
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else if root.Right == nil && root.Left != nil {
			return root.Left
		} else {
			return nil
		}
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
func main() {

}
