package new_queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(6)
	q.EnQueue(2)
	q.EnQueue(3)
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}

func TestCycleArrayQueue(t *testing.T) {
	q := NewCycleArrayQueue(6)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}
