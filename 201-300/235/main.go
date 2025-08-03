package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	tmp := root
	for {
		if p.Val > tmp.Val && q.Val > tmp.Val {
			tmp = tmp.Right
		} else if p.Val < tmp.Val && q.Val < tmp.Val {
			tmp = tmp.Left
		} else {
			return tmp
		}
	}
}
func main() {

}
