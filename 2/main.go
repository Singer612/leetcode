package main

/*
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersSelf(l1 *ListNode, l2 *ListNode) *ListNode {
	//记录一个头结点 ，为了方便最后返回除虚拟节点外的其他值
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	//创建一个链接用来存放结果
	result := head
	//保存当前进位
	carryBit := 0
	if l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := x + y + carryBit
		carryBit = sum / 10
		currentNodeValue := sum % 10
		currentNode := &ListNode{
			Val:  currentNodeValue,
			Next: nil,
		}
		result.Next = currentNode
		//要将指针指到下一个节点
		result = result.Next
		if l1.Next != nil {
			l1 = l1.Next
		}
		if l2.Next != nil {
			l2 = l2.Next
		}
	}
	// 只有最后一位才需要判断是否需要单独存储进位产生的1
	if carryBit == 1 {
		tailNode := &ListNode{
			Val:  carryBit,
			Next: nil,
		}
		result.Next = tailNode
	}
	return head.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 构建一个链表用来存放 l1 和 l2 两个链表相加的结果
	// 其中 dummy 这个节点为虚拟头结点
	dummy := &ListNode{-1, nil}
	// 设置一个进位，初始化为 0
	// 两个个位数相加，进位只能是 1 或者 0
	// 比如 7 + 8 = 15，进位是 1
	// 比如 2 + 3 = 6，没有进位，或者说进位是 0
	carryBit := 0
	// l1 和 l2 有可能为空，所以先默认结果链表从虚拟头结点位置开始
	cur := dummy
	// 同时遍历 l1 和 l2
	// 只要此时 l1 和 l2 两个链表中的任意链表中节点有值，就一直加下去
	// 直到两个链表中的所有节点都遍历完毕为止
	for l1 != nil || l2 != nil {
		// 获取 l1 链表中节点的值
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		// 获取 l2 链表中节点的值
		y := 0
		if l2 != nil {
			y = l2.Val
		}
		// 每一位计算的同时需要考虑上一位的进位情况
		sum := x + y + carryBit
		// 获取当前计算结果的十位数
		carryBit = sum / 10
		// 获取当前计算结果的个位数
		digit := sum % 10
		// 构建一个节点用来存放这个个位数
		digitNode := &ListNode{digit, nil}
		// 把这个节点加入到结果链表中
		cur.Next = digitNode
		// 移动 cur 的位置，观察后面应该存放什么节点
		cur = cur.Next
		// l1 链表中还有节点未遍历完毕就继续遍历下去
		if l1 != nil {
			l1 = l1.Next
		}
		// l2 链表中还有节点未遍历完毕就继续遍历下去
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// 两个链表的尾节点相加之后，有可能产生进位的情况
	// 所以，需要构建一个新的节点用来存放这个进位的结果
	if carryBit == 1 {
		// 构建一个节点用来存放这个进位
		carryBitNode := &ListNode{carryBit, nil}
		// 把这个节点加入到结果链表中
		cur.Next = carryBitNode
	}
	// 最后返回结果链表的头节点就行，即虚拟头结点的下一个节点
	return dummy.Next
}

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	addTwoNumbers(&l1, &l2)
}
