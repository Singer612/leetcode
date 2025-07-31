package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		// false的情况
		if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right == nil {
			return true
		} else if left.Val != right.Val {
			return false
		}
		// 比较左子树的左孩子和右子树的右孩子
		outside := compare(left.Left, right.Right)
		// 比较左子树的右孩子和右子树的左孩子
		inside := compare(left.Right, right.Left)
		return inside && outside
	}
	return compare(root.Left, root.Right)
}

func main() {

}
