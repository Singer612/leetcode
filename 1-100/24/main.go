package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next
		tmp1 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp
		tmp.Next = tmp1
		cur = tmp

	}
	return dummy.Next
}

func main() {

}
