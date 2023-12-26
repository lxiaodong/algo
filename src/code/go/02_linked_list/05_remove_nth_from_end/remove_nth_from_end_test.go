package remove_nth_from_end

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd1(t *testing.T) {
	l := NewLinkedList()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	l.AddAtTail(4)
	l.AddAtTail(5)
	res := l.RemoveNthFromEnd1(1)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func TestRemoveNthFromEnd2(t *testing.T) {
	l := NewLinkedList()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)
	l.AddAtTail(4)
	l.AddAtTail(5)
	res := l.RemoveNthFromEnd2(2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
