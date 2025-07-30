package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func preorder(node *TreeNode, res []int) {
//	if node == nil {
//		return
//	}
//	res = append(res, node.Val)
//	preorder(node.Left, res)
//	preorder(node.Right, res)
//	return
//}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
		return
	}
	preorder(root)
	return res
}

func main() {

}
