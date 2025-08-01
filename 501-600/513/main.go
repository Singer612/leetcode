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

func findBottomLeftValue1(root *TreeNode) int {
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
func findBottomLeftValue(root *TreeNode) int {
	maxDepth, res := 0, 0
	var getHeight func(node *TreeNode, depth int)
	getHeight = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				res = node.Val
			}
			return
		}
		// 从上往下算高度
		if node.Left != nil {
			getHeight(node.Left, depth+1)
		}
		if node.Right != nil {
			getHeight(node.Right, depth+1)
		}
	}
	getHeight(root, 1)
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func buildTree() *TreeNode {
	// 创建节点
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func main() {
	tree := buildTree()
	findBottomLeftValue(tree)
}
