package stack

// 用队列实现栈 https://leetcode.cn/problems/implement-stack-using-queues/

type Stack struct {
	queue []any
}

func NewStack() *Stack {
	return &Stack{
		queue: make([]any, 0),
	}
}

// Push 添加元素
func (s *Stack) Push(v any) {
	s.queue = append(s.queue, v)
}

// Pop 弹出栈顶元素
func (s *Stack) Pop() any {
	n := len(s.queue) - 1
	for n != 0 {
		val := s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, val)
		n--
	}

	val := s.queue[0]
	s.queue = s.queue[1:]
	return val
}

// Empty 判断是否为空
func (s *Stack) Empty() bool {
	return len(s.queue) == 0
}

// Pop 返回栈顶元素
func (s *Stack) Top() any {
	val := s.Pop()
	s.queue = append(s.queue, val)
	return val
}
