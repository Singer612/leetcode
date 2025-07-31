package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
type Queue struct {
	queue []*Node
	size  int
}

func New() *Queue {
	return &Queue{
		queue: make([]*Node, 0),
		size:  0,
	}
}
func (q *Queue) Push(node *Node) {
	q.queue = append(q.queue, node)
	q.size++
}
func (q *Queue) Pop() {
	q.queue = q.queue[1:]
	q.size--
}
func (q *Queue) Top() *Node {
	return q.queue[0]
}
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func connect(root *Node) *Node {
	//res := make([][]int, 0)
	if root == nil {
		return nil
	}
	q := New()
	q.Push(root)
	for !q.IsEmpty() {
		size := q.size
		tmp := make([]int, 0)
		pre := q.Top()
		for i := 0; i < size; i++ {
			node := q.Top()
			if i > 0 {
				pre.Next = node
				pre = pre.Next
			}
			tmp = append(tmp, node.Val)
			q.Pop()
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}

		}
		//res = append(res, tmp)
	}
	return root
}

func main() {

}
