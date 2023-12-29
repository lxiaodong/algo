package cycle_list

import (
	"fmt"
	"testing"
)

func TestCycleList(t *testing.T) {
	list := NewLinkedList()
	list.AddAtIndex(0, 1)
	list.AddAtIndex(1, 2)
	list.AddAtIndex(2, 3)
	list.AddAtIndex(3, 4)
	list.AddAtIndex(4, 5)
	list.AddAtIndex(5, 6)
	list.SetTailNextByIndex(3)
	res := list.DetectCycle()
	fmt.Println(res.Val)
}
