package binary_search

import (
	"fmt"
	"testing"
)

// 测试数组地址是否连续
func TestArray(t *testing.T) {
	arr32 := [10]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i := 0; i < len(arr32); i++ {
		fmt.Printf("%d addr is: %p \n", arr32[i], &arr32[i])
	}
}

// 测试二维数组地址是否连续
func TestTwoDimensionalArray(t *testing.T) {
	arr32 := [2][2]int32{{0, 1}, {2, 3}}

	for i := 0; i < len(arr32); i++ {
		for j := 0; j < len(arr32[i]); j++ {
			fmt.Printf("%d addr is: %p \n", arr32[i][j], &arr32[i][j])
		}
	}
}

// 测试切片参数
func TestSliceParams(t *testing.T) {
	f := func(s []int) int {
		fmt.Printf("s addr:%p\n", &s)
		fmt.Printf("s val addr:%p\n", &s[0])
		return len(s)
	}

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s1 addr:%p\n", &s1)
	fmt.Printf("s1 val addr:%p\n", &s1[0])
	l := f(s1)
	fmt.Println(l)
}

// 测试二分查询左闭右闭
func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	idx := BinarySearch(nums, target)
	fmt.Printf("idx=%d\n", idx)
}

// 测试二分查找左闭右开
func TestBinarySearch1(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	idx := BinarySearch1(nums, target)
	fmt.Printf("idx=%d\n", idx)
}
