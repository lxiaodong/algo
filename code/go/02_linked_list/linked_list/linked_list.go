package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Size int
}

func Constructor() *LinkedList {
	return &LinkedList{
		Head: &ListNode{Val: 0, Next: nil},
		Size: 0,
	}
}

// Get 获取链表中第 index 个节点的值。如果索引无效，则返回-1
func (l *LinkedList) Get(index int) int {
	if index < 0 || index >= l.Size {
		return -1
	}

	cur := l.Head
	for index >= 0 {
		cur = cur.Next
		index--
	}

	return cur.Val
}

// AddAtHead 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (l *LinkedList) AddAtHead(val int) {
	node := &ListNode{Val: val}
	node.Next = l.Head.Next
	l.Head.Next = node
	l.Size++
}

// AddAtTail 将值为 val 的节点追加到链表的最后一个元素
func (l *LinkedList) AddAtTail(val int) {
	node := &ListNode{Val: val}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
	l.Size++
}

// AddAtIndex 在链表中的第 index 个节点之前添加值为 val  的节点。
// 如果 index 等于链表的长度，则该节点将附加到链表的末尾。
// 如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (l *LinkedList) AddAtIndex(index int, val int) {
	if index > l.Size {
		return
	}

	if index < 0 {
		index = 0
	}

	current := l.Head
	for index > 0 {
		current = current.Next
		index--
	}

	node := &ListNode{Val: val}
	node.Next = current.Next
	current.Next = node
	l.Size++
}

// DeleteAtIndex 如果索引 index 有效，则删除链表中的第 index 个节点
func (l *LinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Size {
		return
	}

	slow := l.Head
	fast := l.Head.Next

	for index > 0 {
		fast = fast.Next
		slow = slow.Next
		index--
	}
	slow.Next = fast.Next
	l.Size--
}

func (l *LinkedList) Print() {
	current := l.Head
	for current.Next != nil {
		fmt.Println(current.Next.Val)
		current = current.Next
	}
}
