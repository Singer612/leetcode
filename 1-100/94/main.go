package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

type Stack struct {
	stack []*TreeNode
}

func New() *Stack {
	return &Stack{
		stack: make([]*TreeNode, 0),
	}
}
func (s *Stack) Pop() {
	if !s.IsEmpty() {
		s.stack = s.stack[:len(s.stack)-1]
	}
	return
}
func (s *Stack) Push(node *TreeNode) {
	s.stack = append(s.stack, node)

}
func (s *Stack) Top() *TreeNode {
	if !s.IsEmpty() {
		return s.stack[len(s.stack)-1]
	}
	return nil
}
func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	s := New()
	cur := root
	for cur != nil || len(s.stack) != 0 {
		// 还有1个终止条件，当len(s.stack) != 0时，cur可能为空
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		if cur == nil {
			cur = s.Top()
			s.Pop()
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

func main() {

}
