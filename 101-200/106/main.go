package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	tmp := 0
	for i, v := range inorder {
		if rootVal == v {
			tmp = i
			break
		}
	}
	root := &TreeNode{
		Val: rootVal,
	}
	left := inorder[:tmp]
	right := inorder[tmp+1:]
	root.Right = buildTree(right, postorder[:len(postorder)-1])
	root.Left = buildTree(left, postorder[:len(postorder)-1-len(right)])
	return root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	buildTree(inorder, postorder)
}
