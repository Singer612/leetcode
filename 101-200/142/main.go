package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		// 快指针走2步，慢指针走1步
		fast = fast.Next.Next
		slow = slow.Next
		// 如果找到环，就去找环的入口
		if fast == slow {
			index := slow
			pre := head
			for pre != index {
				pre = pre.Next
				index = index.Next
			}
			return pre
		}
	}
	return nil
}

func main() {

}
