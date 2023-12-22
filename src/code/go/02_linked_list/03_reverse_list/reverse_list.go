package reverse_list

import "fmt"

// 反转链表 https://leetcode.cn/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: &ListNode{Val: 0},
		Size: 0,
	}
}

func (l *LinkedList) AddAtTail(val int) {
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	node := &ListNode{Val: val}
	current.Next = node
	l.Size++
}

// Reverse1 双指针方方法
// 1 -> 2 -> 3 -> 4 -> 5 -> nil
// nil <- 1 <- 2 <- 3 <- 4 <- 5
func (l *LinkedList) Reverse1() *ListNode {
	current := l.Head.Next
	var pre *ListNode
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre
}

// Reverse2 递归法
// 1 -> 2 -> 3 -> 4 -> 5 -> nil
// nil <- 1 <- 2 <- 3 <- 4 <- 5
func (l *LinkedList) Reverse2() *ListNode {
	return help(nil, l.Head.Next)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}

	next := head.Next
	head.Next = pre
	return help(head, next)
}

func Print(l *ListNode) {
	head := &ListNode{Val: 0}
	head.Next = l
	current := head
	for current.Next != nil {
		fmt.Println(current.Next.Val)
		current = current.Next
	}
}
