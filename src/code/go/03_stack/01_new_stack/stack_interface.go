package new_stack

type Stack interface {
	Push(any)
	Pop() any
	Top() any
	Flush()
	IsEmpty() bool
}
