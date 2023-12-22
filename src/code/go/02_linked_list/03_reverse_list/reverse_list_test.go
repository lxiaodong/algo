package reverse_list

import "testing"

// 测试双指针方法
func TestReverse1(t *testing.T) {
	list := NewLinkedList()
	list.AddAtTail(1)
	list.AddAtTail(2)
	list.AddAtTail(3)
	list.AddAtTail(4)
	// Print(list.Head.Next)

	l1 := list.Reverse1()
	Print(l1)
}

// 测试递归方法
func TestReverse2(t *testing.T) {
	list := NewLinkedList()
	list.AddAtTail(1)
	list.AddAtTail(2)
	list.AddAtTail(3)
	list.AddAtTail(4)
	// Print(list.Head.Next)

	l2 := list.Reverse2()
	Print(l2)
}
