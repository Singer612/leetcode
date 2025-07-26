package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: nil,
		Size: 0,
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.Size || l.Head == nil {
		return -1
	}
	cur := l.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{
		Val:  val,
		Next: l.Head,
	}
	l.Head = node
	l.Size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = node
		l.Size++
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	l.Size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	}
	if index > l.Size {
		return
	}
	if index == 0 {
		l.AddAtHead(val)
		return
	}
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	node := &ListNode{
		Val:  val,
		Next: cur.Next,
	}
	cur.Next = node
	l.Size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Size || l.Head == nil {
		return
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Size--
		return
	}
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	l.Size--
}

func main() {

}
