package _1_array

import (
	"fmt"
	"testing"
)

// 测试最小子数组长度
func TestMinSubArrayLen1(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	res := MinSubArrayLen1(nums, target)
	fmt.Println(res)
}

// 测试最小子数组长度
func TestMinSubArrayLen2(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	res := MinSubArrayLen2(nums, target)
	fmt.Println(res)
}
