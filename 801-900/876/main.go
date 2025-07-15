package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(length int) *ListNode {
	if length <= 0 {
		return nil
	}

	head := &ListNode{Val: 1}
	current := head

	for i := 2; i <= length; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
	}

	return head
}

func middleNode(head *ListNode) *ListNode {
	NodeMap := make(map[int]*ListNode)
	index := 0
	dummy := &ListNode{
		Next: head,
	}
	for tmp := dummy; tmp.Next != nil; tmp = tmp.Next {
		index++
		NodeMap[index] = tmp.Next // 修正：存储当前节点而非前一个节点

	}
	result := index / 2
	fmt.Printf("中间节点索引: %d\n", result)
	fmt.Printf("节点总数: %d\n", index)

	// 调试：打印所有节点值
	for i := 1; i <= index; i++ {
		fmt.Printf("索引 %d: 值 %d\n", i, NodeMap[i].Val)
	}

	return NodeMap[result]
}

func main() {
	list := createLinkedList(6)
	result := middleNode(list)
	fmt.Println(result)
}
