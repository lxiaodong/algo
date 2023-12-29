package cycle_list

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

	node := &ListNode{Val: val, Next: current.Next}
	current.Next = node
	l.Size++
}

func (l *LinkedList) SetTailNextByIndex(index int) {
	tail := l.Head
	for tail.Next != nil {
		tail = tail.Next
	}

	current := l.Head
	for index >= 0 {
		current = current.Next
		index--
	}

	tail.Next = current
}

// 检查是否有环并且检查环的入口
func (l *LinkedList) DetectCycle() *ListNode {
	current := l.Head
	fast := current
	slow := current

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 快慢指针相遇，此时从head 和 相遇点，同时查找直至相遇
		if slow == fast {
			index1 := fast
			index2 := current
			for index1 != index2 {
				index2 = index2.Next
				index1 = index1.Next
			}
			return index2
		}
	}

	return nil
}
