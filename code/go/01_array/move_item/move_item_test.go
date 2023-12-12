package _1_array

import (
	"fmt"
	"testing"
)

// 测试粗暴移除数组元素
func TestMoveItem1(t *testing.T) {
	// nums := []int{3,2,2,3}
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	len := MoveItem1(nums, target)
	fmt.Println(len)
}

// 测试双指针移除数组元素
func TestMoveItem2(t *testing.T) {
	// nums := []int{3,2,2,3}
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	len := MoveItem2(nums, target)
	fmt.Println(len)
}

// 测试双指针移除数组元素
func TestMoveItem3(t *testing.T) {
	// nums := []int{3,2,2,3}
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	len := MoveItem3(nums, target)
	fmt.Println(len)
}
