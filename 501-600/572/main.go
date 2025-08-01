package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var compare func(root *TreeNode, subRoot *TreeNode) bool
	compare = func(root *TreeNode, subRoot *TreeNode) bool {
		if root == nil && subRoot == nil {
			return true
		} else if root == nil && subRoot != nil {
			return false
		} else if root != nil && subRoot == nil {
			return false
		} else if root.Val != subRoot.Val {
			return false
		}
		return compare(root.Left, subRoot.Left) && compare(root.Right, subRoot.Right)
	}
	// 遍历root树
	var traversal func(node *TreeNode) bool
	traversal = func(node *TreeNode) bool {

		//如果节点为空，res=false，那就是找不到了
		if node == nil {
			return false
		}
		if compare(node, subRoot) {
			return true
		}
		return traversal(node.Left) || traversal(node.Right)
	}
	return traversal(root)
}

func main() {

}
