package main

type MyCircularDeque struct {
	front, rear int   // 队头队尾指针
	elements    []int // 存储队列元素的数组
}

func Constructor(k int) MyCircularDeque {
	// 创建一个比当前队列大一个的数组
	return MyCircularDeque{elements: make([]int, k+1)}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	// 队列满了，无法插入数据
	if q.IsFull() {
		return false
	}
	q.front = (q.front - 1 + len(q.elements)) % len(q.elements)
	q.elements[q.front] = value
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	// 队列满了，无法插入数据
	if q.IsFull() {
		return false
	}
	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	// 队列空了，无法删除
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % len(q.elements)
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	// 队列空了，无法删除
	if q.IsEmpty() {
		return false
	}
	q.rear = ((q.rear - 1) + len(q.elements)) % len(q.elements)
	return true
}

func (q *MyCircularDeque) GetFront() int {
	// 队列空了，无法获取
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q *MyCircularDeque) GetRear() int {
	// 队列空了，无法获取
	if q.IsEmpty() {
		return -1
	}
	return q.elements[((q.rear-1)+len(q.elements))%len(q.elements)]
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.rear == q.front
}

func (q *MyCircularDeque) IsFull() bool {
	return (q.rear+1)%len(q.elements) == q.front
}

func main() {

}
