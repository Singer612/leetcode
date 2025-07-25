package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 先处理头结点
	cur := head
	// 找到反转后的下一个节点
	var pre *ListNode
	// 遍历链表到尾
	for cur != nil {
		// 保存下来，用来链接下一个节点
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {

}
