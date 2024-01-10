package valid_parentheses

// 有效的括号 https://leetcode.cn/problems/valid-parentheses/
type Stack struct {
	data []any
	top  int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]any, 32),
		top:  -1,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Push(v any) bool {
	if s.top == cap(s.data)-1 {
		return false
	}
	s.top++
	s.data[s.top] = v
	return true
}

func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top--
	return v
}

func (s *Stack) Top() any {
	return s.data[s.top]
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	hash := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := NewStack()

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack.Push(s[i])
		} else if !stack.IsEmpty() && hash[s[i]] == stack.Top() {
			stack.Pop()
		} else {
			return false
		}
	}
	return stack.IsEmpty()
}
