package top_k_frequent

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(TopKFrequent(nums, k))
}
