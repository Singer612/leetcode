package main

func oddEvenList(head *ListNode) *ListNode {
	// 保留偶数链表头结点

	head1, head2 := head, head.Next
	second := head2
	//奇数偶数链表都开始遍历
	for head1 != nil || head2 != nil {
		head1 = head2.Next
		head2 = head1.Next
	}
	head1.Next = second
	return head1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
