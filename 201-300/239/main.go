package main

// 封装单调队列，主要有三个操作
// push:如果当前元素比队尾元素大，那么将队尾元素排出，直到当前元素进队列
// pop:排出队列头元素，也就是最大的元素
// front:获取队列头最大的元素

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}
func (q *MyQueue) Push(val int) {
	for len(q.queue) != 0 && q.queue[len(q.queue)-1] < val {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, val)

}
func (q *MyQueue) Pop(val int) {
	if q.queue[0] == val {
		q.queue = q.queue[1:]
	}
}
func (q *MyQueue) Front() int {
	return q.queue[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	q := NewMyQueue()
	res := make([]int, 0)
	// 将第一个滑动窗口的k个元素入队
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	// 找到第一个滑动窗口的最大值
	res = append(res, q.Front())

	for i := k; i < len(nums); i++ {
		// 滑动窗口向右移，pop最左边的元素（如果队列里有，没有的话不作操作）
		q.Pop(nums[i-k])
		// 加入新的元素
		q.Push(nums[i])
		// 获取结果
		res = append(res, q.Front())
	}
	return res
}

func main() {

}
