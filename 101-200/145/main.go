package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
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
func postorderTraversal(root *TreeNode) []int {
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
		if node.Left != nil {
			s.Push(node.Left)
		}
		if node.Right != nil {
			s.Push(node.Right)
		}
	}
	return reverse(res)
}

func reverse(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

func main() {

}
