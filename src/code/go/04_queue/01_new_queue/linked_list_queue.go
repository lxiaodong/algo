package new_queue

type Node struct {
	Val  any
	Next *Node
}

var _ Queue = (*LinkedListQueue)(nil)

type LinkedListQueue struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// EnQueue 入队
func (q *LinkedListQueue) EnQueue(v any) bool {
	node := &Node{Val: v}
	if q.tail == nil {
		q.tail = node
		q.head = node
	} else {
		q.tail.Next = node
		q.tail = node
	}
	q.length++
	return true
}

// DeQueue 出对
func (q *LinkedListQueue) DeQueue() any {
	if q.head == nil {
		return nil
	}

	v := q.head.Val
	q.head = q.head.Next
	q.length--
	return v
}
