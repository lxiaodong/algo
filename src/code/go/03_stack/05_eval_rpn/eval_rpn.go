package eval_rpn

import (
	"strconv"
)

// 逆波兰表达式求值 https://leetcode.cn/problems/evaluate-reverse-polish-notation/

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

// Top 返回栈顶元素
func (s *Stack) Top() any {
	return s.data[s.top]
}

// Push 添加数据
func (s *Stack) Push(v any) bool {
	if cap(s.data)-1 == s.top {
		return false
	}

	s.top++
	s.data[s.top] = v
	return true
}

// Pop 弹出数据
func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}

	v := s.data[s.top]
	s.top--
	return v
}

func EvalRPN1(tokens []string) int {
	stack := NewStack()
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack.Push(val)
		} else {
			num1, num2 := stack.Pop().(int), stack.Pop().(int)
			val := 0
			switch token {
			case "+":
				val = num1 + num2
			case "-":
				val = num1 - num2
			case "*":
				val = num1 * num2
			case "/":
				val = num1 / num2
			}
			stack.Push(val)
		}
	}
	return stack.Top().(int)
}

func EvalRPN2(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			val := 0
			switch token {
			case "+":
				val = num1 + num2
			case "-":
				val = num1 - num2
			case "*":
				val = num1 * num2
			case "/":
				val = num1 / num2
			}
			stack = append(stack, val)
		}
	}
	return stack[0]
}
