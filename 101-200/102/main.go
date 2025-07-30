package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Queue struct {
	queue []*TreeNode
	size  int
}

func New() *Queue {
	return &Queue{
		queue: make([]*TreeNode, 0),
		size:  0,
	}
}
func (q *Queue) Push(node *TreeNode) {
	q.queue = append(q.queue, node)
	q.size++
}
func (q *Queue) Pop() {
	q.queue = q.queue[1:]
	q.size--
}
func (q *Queue) Top() *TreeNode {
	return q.queue[0]
}
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := New()
	q.Push(root)
	for !q.IsEmpty() {
		size := q.size
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := q.Top()
			tmp = append(tmp, node.Val)
			q.Pop()
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

func main() {

}
