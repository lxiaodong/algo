package new_stack

var _ Stack = (*ArrayStack)(nil)

type ArrayStack struct {
	data []any
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]any, 32),
		top:  -1,
	}
}

// Push 入栈
func (s *ArrayStack) Push(v any) {
	if s.top == cap(s.data)-1 {
		return
	}

	s.top += 1
	s.data[s.top] = v
}

// Pop 出栈
func (a *ArrayStack) Pop() any {
	if a.IsEmpty() {
		return nil
	}

	v := a.data[a.top]
	a.top--
	return v
}

// IsEmpty 是否为空
func (a *ArrayStack) IsEmpty() bool {
	return len(a.data) == 0 && a.top == -1
}

// Flush 清空
func (a *ArrayStack) Flush() {
	a.top = -1
}

// Top 置顶元素
func (a *ArrayStack) Top() any {
	if a.IsEmpty() {
		return nil
	}

	return a.data[a.top]
}
