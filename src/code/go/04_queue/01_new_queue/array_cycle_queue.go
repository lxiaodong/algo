package new_queue

var _ Queue = (*CycleArrayQueue)(nil)

type CycleArrayQueue struct {
	data     []any
	capacity int
	head     int
	tail     int
}

func NewCycleArrayQueue(capacity int) *CycleArrayQueue {
	return &CycleArrayQueue{
		data:     make([]any, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

// DeQueue 出对
func (q *CycleArrayQueue) DeQueue() any {
	if q.tail == q.head {
		return nil
	}
	v := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

// EnQueue 入队
func (q *CycleArrayQueue) EnQueue(v any) bool {
	if (q.tail+1)%q.capacity == q.head {
		return false
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}
