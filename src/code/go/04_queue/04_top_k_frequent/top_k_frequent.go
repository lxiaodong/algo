package top_k_frequent

import (
	"container/heap"
)

// 前 K 个高频元素 https://leetcode.cn/problems/top-k-frequent-elements/

// 小顶堆
type IHeap [][2]int

// Len 返回长度
func (h IHeap) Len() int {
	return len(h)
}

// Less 比较
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

// Swap 交换
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 添加元素
func (h *IHeap) Push(v any) {
	*h = append(*h, v.([2]int))
}

// Pop 弹出元素
func (h *IHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 方法1： 小顶堆
func TopKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}

	// 记录每个元素出现的次数
	for _, num := range nums {
		map_num[num]++
	}

	h := &IHeap{}
	heap.Init(h)

	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	//按顺序返回堆中的元素

	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}
