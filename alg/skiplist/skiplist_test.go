package skiplist

import (
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
	sl.DebugPrintSkipList()

}
