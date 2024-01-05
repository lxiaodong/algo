package new_stack

type node struct {
	val  any
	next *node
}

var _ Stack = (*LinkedListStack)(nil)

type LinkedListStack struct {
	top *node
}

func NewLinkedStack() *LinkedListStack {
	return &LinkedListStack{}
}

// Push 入栈
func (l *LinkedListStack) Push(val any) {
	l.top = &node{val: val, next: l.top}
}

// Pop 出栈
func (l *LinkedListStack) Pop() any {
	if l.IsEmpty() {
		return nil
	}

	node := l.top
	l.top = l.top.next
	return node.val
}

// IsEmpty 是否为空
func (l *LinkedListStack) IsEmpty() bool {
	return l.top == nil
}

// Top 置顶元素
func (l *LinkedListStack) Top() any {
	if l.IsEmpty() {
		return nil
	}
	return l.top.val
}

// Flush 清空
func (l *LinkedListStack) Flush() {
	l.top = nil
}
