package swap_pairs

// 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/

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

	node := &ListNode{Val: val}
	current.Next = node
	l.Size++
}

func (l *LinkedList) SwapPairs() *ListNode {
	current := l.Head
	for current.Next != nil && current.Next.Next != nil {
		nn := current.Next.Next
		n := current.Next

		n.Next = nn.Next
		nn.Next = n
		current.Next = nn

		current = current.Next.Next
	}

	return l.Head.Next
}
