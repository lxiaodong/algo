package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := Constructor()     // 初始化链表
	list.AddAtHead(100)       // 在头部添加元素
	list.AddAtTail(242)       // 在尾部添加元素
	list.AddAtTail(777)       // 在尾部添加元素
	list.AddAtIndex(1, 99999) // 在指定位置添加元素
	list.Print()              // 打印链
	fmt.Println(list.Get(1))
}
