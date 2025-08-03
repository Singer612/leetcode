package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	tmp := &TreeNode{
		Val: val,
	}
	if root == nil {
		return tmp
	}
	var insert func(node *TreeNode, val int, flg bool)
	insert = func(node *TreeNode, val int, flg bool) {
		if flg {
			return
		}
		if val > node.Val {
			if node.Right == nil {
				node.Right = tmp
				insert(node.Right, val, true)
			} else {
				insert(node.Right, val, false)
			}

		}
		if val < node.Val {
			if node.Left == nil {
				node.Left = tmp
				insert(node.Left, val, true)
			} else {
				insert(node.Left, val, false)
			}
		}
	}
	insert(root, val, false)
	return root
}
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
func main() {

}
