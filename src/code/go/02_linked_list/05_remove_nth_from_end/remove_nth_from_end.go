package remove_nth_from_end

// 删除链表的倒数第N个节点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

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
		Head: &ListNode{
			Val: 0,
		},
		Size: 0,
	}
}

func (l *LinkedList) AddAtTail(val int) {
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &ListNode{Val: val}
	l.Size++
}

func (l *LinkedList) RemoveNthFromEnd1(n int) *ListNode {
	current := l.Head
	idx := l.Size - n

	for idx > 0 {
		current = current.Next
		idx--
	}

	current.Next = current.Next.Next
	l.Size--

	return l.Head.Next
}

func (l *LinkedList) RemoveNthFromEnd2(n int) *ListNode {
	fast := l.Head
	slow := l.Head

	for n > 0 && fast.Next != nil {
		fast = fast.Next
		n--
	}

	fast = fast.Next
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	l.Size--

	return l.Head.Next
}
