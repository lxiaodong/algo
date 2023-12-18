package remove_element

import (
	"fmt"
	"testing"
)

// 测试移除链表元素
func TestRemoveElement(t *testing.T) {
	list := NewList(1, 2, 6, 3, 4, 5, 6, 8, 7)
	resp := RemoveElement(list, 6)
	for resp != nil {
		fmt.Println(resp.Val)
		resp = resp.Next
	}
}
