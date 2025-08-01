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

// 这种只传1个节点的逻辑并不正确
// 因为没有办法同步判断左右子树是否相等，无法直接控制递归时比较的节点对
// 所以需要自己手动控制，也就是说传左右节点
func isSymmetric1(root *TreeNode) bool {
	var compare func(node *TreeNode) bool
	compare = func(node *TreeNode) bool {
		// false的情况
		if node.Left == nil && node.Right != nil {
			return false
		} else if node.Left != nil && node.Right == nil {
			return false
		} else if node.Left == nil && node.Right == nil {
			return true
		} else if node.Left.Val != node.Right.Val {
			return false
		}
		// 比较左子树的左孩子和右子树的右孩子
		outside := compare(node.Left)
		// 比较左子树的右孩子和右子树的左孩子
		inside := compare(node.Right)
		return inside && outside
	}
	return compare(root)
}

func main() {

}
