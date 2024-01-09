package queue

// 用栈实现队列 https://leetcode.cn/problems/implement-queue-using-stacks/

type Stack struct {
	data []any
	top  int
}

type Queue struct {
	In  *Stack
	Out *Stack
}

func NewQueue() *Queue {
	return &Queue{
		In: &Stack{
			data: make([]any, 32),
			top:  -1,
		},
		Out: &Stack{
			data: make([]any, 32),
			top:  -1,
		},
	}
}

// Push 添加元素
func (q *Queue) Push(v any) bool {
	if cap(q.In.data)-1 == q.In.top {
		return false
	}
	q.In.top++
	q.In.data[q.In.top] = v
	return true
}

// Pop 弹出最上边元素
func (q *Queue) Pop() any {
	if q.Out.top == -1 && q.In.top == -1 {
		return nil
	}

	if q.Out.top == -1 {
		for q.In.top != -1 {
			v := q.In.data[q.In.top]
			q.In.top--
			q.Out.top++
			q.Out.data[q.Out.top] = v
		}
	}
	popV := q.Out.data[q.Out.top]
	q.Out.top--
	return popV
}
