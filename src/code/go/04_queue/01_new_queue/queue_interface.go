package new_queue

type Queue interface {
	EnQueue(any) bool
	DeQueue() any
}
