package new_queue

var _ Queue = (*ArrayQueue)(nil)

type ArrayQueue struct {
	data     []any
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		data:     make([]any, capacity),
		head:     0,
		tail:     0,
		capacity: capacity,
	}
}

// EnQueue 入队
func (q *ArrayQueue) EnQueue(v any) bool {
	if q.tail == q.capacity {
		if q.head == 0 {
			return false
		}
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
		}
		q.tail -= q.head
		q.head = 0
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

// DeQueue 出对
func (q *ArrayQueue) DeQueue() any {
	if q.head == q.tail {
		return nil
	}

	v := q.data[q.head]
	q.head++

	return v
}
