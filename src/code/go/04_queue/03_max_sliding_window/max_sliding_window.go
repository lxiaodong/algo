package max_sliding_window

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

// Front 返回第一个元素
func (q *Queue) Front() int {
	return q.data[0]
}

// Back 返回最后一个元素
func (q *Queue) Back() int {
	return q.data[len(q.data)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Push 增加元素
// 如果push的元素value大于入口元素的数值
// 那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
func (q *Queue) Push(v int) {
	for !q.IsEmpty() && v > q.Back() {
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, v)
}

// Pop 如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
func (q *Queue) Pop(v int) {
	if !q.IsEmpty() && v == q.Front() {
		q.data = q.data[1:]
	}
}

func MaxSlidingWindow(nums []int, k int) []int {
	queue := NewQueue()
	length := len(nums)
	res := make([]int, 0)

	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}

	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}

	return res
}
