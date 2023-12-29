package intersection_node

// 链表相交 https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

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

func (l *LinkedList) AddListAtTail(list *ListNode) {
	head := l.Head
	for head.Next != nil {
		head = head.Next
	}

	ll := list
	listLen := 0
	for ll != nil {
		ll = ll.Next
		listLen++
	}

	head.Next = list
	l.Size += listLen
}

// 1->2->3->10->11->12->4->11->12
// 4->11->12->1->2->3->10->11->12
// 双指针方法
func GetIntersectionNode1(a, b *ListNode) *ListNode {
	h1, h2 := a, b
	for h1 != h2 {
		if h1 != nil {
			h1 = h1.Next
		} else {
			h1 = b
		}
		if h2 != nil {
			h2 = h2.Next
		} else {
			h2 = a
		}
	}

	return h1
}

func GetIntersectionNode2(a, b *ListNode) *ListNode {
	// 获取各个链表的长度
	l1, l2 := a, b

	len1 := 0
	for l1 != nil {
		l1 = l1.Next
		len1++
	}

	len2 := 0
	for l2 != nil {
		l2 = l2.Next
		len2++
	}

	var step int
	var fast, slow *ListNode

	if len1 > len2 {
		fast, slow = a, b
		step = len1 - len2
	} else {
		fast, slow = b, a
		step = len2 - len1
	}

	for step > 0 {
		fast = fast.Next
		step--
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast

}
