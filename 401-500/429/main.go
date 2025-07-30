package main

type Node struct {
	Val      int
	Children []*Node
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

func levelOrder(root *Node) [][]int {
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
			for _, v := range node.Children {
				if v != nil {
					q.Push(v)
				}
			}
		}
		res = append(res, tmp)
	}
	return res
}

func main() {

}
