package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Queue 实现队列
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
func (q *Queue) Pop() *TreeNode {
	if !q.IsEmpty() {
		tmp := q.queue[0]
		q.queue = q.queue[1:]
		q.size--
		return tmp
	}
	return nil
}
func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	q := New()
	q.Push(root)
	for !q.IsEmpty() {
		size := q.size
		for i := 0; i < size; i++ {
			tmp := q.Pop()
			if i == 0 {
				res = tmp.Val
			}
			if tmp.Left != nil {
				q.Push(tmp.Left)
			}
			if tmp.Right != nil {
				q.Push(tmp.Right)
			}
		}
	}
	return res
}
func main() {

}
