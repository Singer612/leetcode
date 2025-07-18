package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	// 创建头结点
	head := &ListNode{Val: nums[0]}
	current := head
	// 遍历数组，依次创建后续节点
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -1,
	}
	cur := dummy
	tens := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + tens
		tmp := &ListNode{}
		if sum > 9 {
			ones := sum % 10
			tens = (sum / 10) % 10
			tmp.Val = ones
		} else {
			tmp.Val = sum
			tens = 0
		}
		cur.Next = tmp
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	var l *ListNode
	if l1 != nil {
		l = l1
	}
	if l2 != nil {
		l = l2
	}
	for l != nil {
		sum := l.Val + tens
		tmp := &ListNode{}
		if sum > 9 {
			ones := sum % 10
			tens = (sum / 10) % 10
			tmp.Val = ones
		} else {
			tmp.Val = sum
			tens = 0
		}
		cur.Next = tmp
		cur = cur.Next
		l = l.Next
	}
	if tens != 0 {
		tmp := &ListNode{
			Val: tens,
		}
		cur.Next = tmp
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	l1 := createLinkedList([]int{2, 4, 3})
	// 创建 l2 = [5,6,4]
	l2 := createLinkedList([]int{5, 6, 4})
	result := addTwoNumbers(l1, l2)
	fmt.Println(result)
}
