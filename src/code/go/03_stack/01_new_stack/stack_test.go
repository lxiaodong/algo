package new_stack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func TestLinkedListStack(t *testing.T) {
	l := NewLinkedStack()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	fmt.Println(l.Pop())
	fmt.Println(l.Top())
	fmt.Println(l.Pop())
}
