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

func preorderTraversal1(root *TreeNode) []int {
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

// 迭代，使用栈
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	s := New()
	s.Push(root)
	for !s.IsEmpty() {
		node := s.Top()
		res = append(res, node.Val)
		s.Pop()
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
	return res
}

func main() {

}
