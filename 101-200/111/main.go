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

func minDepth1(root *TreeNode) int {
	res := make([][]int, 0)
	depth := 0
	if root == nil {
		return 0
	}
	q := New()
	q.Push(root)
	for !q.IsEmpty() {
		depth++
		size := q.size
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := q.Top()
			tmp = append(tmp, node.Val)
			q.Pop()
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		res = append(res, tmp)
	}
	return depth
}

// 后序遍历
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x > 0 && y > 0 {
		if x < y {
			return x
		}
		return y
	} else if x == 0 {
		return y
	} else {
		return x
	}
}

func main() {

}
