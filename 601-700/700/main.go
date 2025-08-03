package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	var res *TreeNode
	var search func(root *TreeNode, val int)
	search = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		if root.Val > val {
			search(root.Left, val)
		} else if root.Val < val {
			search(root.Right, val)
		} else {
			res = root
			return
		}
	}
	search(root, val)
	return res
}

func main() {

}
