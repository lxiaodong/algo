package remove_element

// 移除链表元素 https://leetcode.cn/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(list ...int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: nil}
	head := dummyHead
	for _, v := range list {
		n := ListNode{Val: v}
		head.Next = &n
		head = head.Next
	}
	return dummyHead.Next
}

func RemoveElement(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	current := dummyHead
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummyHead.Next
}
