package remove_duplicates

// 删除字符串中的所有相邻重复项 https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/

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

// IsEmpty 判断是否为空
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

// Push 增加元素
func (s *Stack) Push(v any) bool {
	if cap(s.data)-1 == s.top {
		return false
	}

	s.top++
	s.data[s.top] = v
	return true
}

// Pop 弹出元素
func (s *Stack) Pop() any {
	v := s.data[s.top]
	s.top--
	return v
}

// Top 返回栈顶元素
func (s *Stack) Top() any {
	return s.data[s.top]
}

// RemoveDuplicates1 结合自定义栈
func RemoveDuplicates1(s string) string {
	if s == "" {
		return s
	}

	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if stack.IsEmpty() || s[i] != stack.Top() {
			stack.Push(s[i])
		} else {
			stack.Pop()
		}
	}

	str := ""
	for !stack.IsEmpty() {
		v := stack.Pop()
		str = string(v.(byte)) + str
	}
	return str
}

// RemoveDuplicates2 结合切片（也是栈）
func RemoveDuplicates2(s string) string {
	if s == "" {
		return s
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] != stack[len(stack)-1] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
