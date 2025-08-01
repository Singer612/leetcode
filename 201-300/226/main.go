package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var invert func(node *TreeNode)
	invert = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		invert(node.Left)
		invert(node.Right)
	}
	invert(root)
	//if root == nil {
	//	return nil
	//}
	//root.Left, root.Right = root.Right, root.Left
	//invertTree(root.Left)
	//invertTree(root.Right)
	return root
}

func main() {

}
