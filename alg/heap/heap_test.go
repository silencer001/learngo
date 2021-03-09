package heap

import (
	"fmt"
	"testing"
)

func TestBuildHeap(t *testing.T) {
	a := []int{10, 3, 4, 5, 6, 1, 13, 11, 12, 8, 7}
	h := NewHeapCopy(a)
	h.BuildHeap()
	fmt.Println(h.s)
	h.Push(2)
	fmt.Println(h.s)
	fmt.Println(h.PopTop())
	fmt.Println(h.s)
	fmt.Println(h.PopTop())
	fmt.Println(h.s)
	fmt.Println(h.PopTop())
	fmt.Println(h.s)
}
