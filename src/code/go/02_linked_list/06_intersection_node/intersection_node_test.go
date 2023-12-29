package intersection_node

import (
	"fmt"
	"testing"
)

func TestIntersectionNode(t *testing.T) {
	listCommon := NewLinkedList()
	listCommon.AddAtTail(12)
	listCommon.AddAtTail(11)

	list1 := NewLinkedList()
	list1.AddAtTail(1)
	list1.AddAtTail(2)
	list1.AddAtTail(3)
	list1.AddAtTail(4)
	list1.AddListAtTail(listCommon.Head.Next)

	list2 := NewLinkedList()
	list2.AddAtTail(8)
	list2.AddAtTail(9)
	list2.AddListAtTail(listCommon.Head.Next)

	res := GetIntersectionNode1(list1.Head.Next, list2.Head.Next)
	fmt.Println(res.Val)
}
