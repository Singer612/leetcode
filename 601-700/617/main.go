package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 != nil {
		return root2
	}
	if root1 != nil && root2 == nil {
		return root1
	}
	if root1 == nil && root2 == nil {
		return nil
	}
	var merge func(root1 *TreeNode, root2 *TreeNode)
	merge = func(root1 *TreeNode, root2 *TreeNode) {
		if root2 == nil || root1 == nil {
			return
		}
		if root1 != nil && root2 != nil {
			root1.Val = root1.Val + root2.Val
		}
		// 做平移逻辑
		if root1.Left == nil && root2.Left != nil {
			root1.Left = root2.Left
			merge(root1.Left, nil)
		} else {
			merge(root1.Left, root2.Left)
		}
		if root1.Right == nil && root2.Right != nil {
			root1.Right = root2.Right
			merge(root1.Right, nil)
		} else {
			merge(root1.Right, root2.Right)
		}
		return
	}
	merge(root1, root2)
	return root1
}

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func main() {

}
