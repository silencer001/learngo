package skiplist

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	sl := NewSkipList()
	sl.SkipListInsert(1, "a")
	sl.SkipListInsert(2, "b")
	sl.SkipListInsert(10, "c")
	sl.SkipListInsert(2, "d")
	sl.SkipListInsert(3, "e")
	sl.SkipListInsert(5, "f")
	sl.SkipListInsert(4, "g")
	sl.SkipListInsert(11, "h")
	sl.SkipListInsert(18, "i")
	sl.SkipListInsert(15, "j")
	sl.SkipListInsert(30, "k")
	sl.SkipListInsert(9, "l")
	sl.DebugPrintSkipList()

}
func TestRandomLevel(t *testing.T) {
	var count int
	for i := 0; i < 20; i++ {
		if randomLevel() > 1 {
			count++
		}
	}
	fmt.Println(count)
}

func TestSkipListFindNode(t *testing.T) {
	sl := NewSkipList()
	sl.SkipListInsert(1, "a")
	sl.SkipListInsert(2, "b")
	sl.SkipListInsert(10, "c")
	sl.SkipListInsert(2, "d")
	sl.SkipListInsert(3, "e")
	sl.SkipListInsert(5, "f")
	sl.SkipListInsert(4, "g")
	sl.SkipListInsert(11, "h")
	sl.SkipListInsert(18, "i")
	sl.SkipListInsert(15, "j")
	sl.SkipListInsert(30, "k")
	sl.SkipListInsert(9, "l")
	sl.DebugPrintSkipList()
	if node := sl.SkipListFindNode(2, "d"); node != nil {
		fmt.Println("found it!")
	} else {
		fmt.Println("not found")
	}
}
