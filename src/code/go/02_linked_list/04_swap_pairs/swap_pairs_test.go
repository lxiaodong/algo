package swap_pairs

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	list := NewLinkedList()
	list.AddAtTail(1)
	list.AddAtTail(2)
	list.AddAtTail(3)
	list.AddAtTail(4)
	list.AddAtTail(5)
	res := list.SwapPairs()
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
