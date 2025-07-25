package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	slow := dummy
	for fast := head; fast != nil; fast = fast.Next {
		if fast.Val != val {
			slow.Next = fast
			slow = slow.Next
		} else {
			slow.Next = nil
		}
	}
	return dummy.Next
}

func main() {

}
