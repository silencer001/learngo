package lru

import (
	"fmt"
	"testing"
	"time"
)

func TestLRU(t *testing.T) {
	li := Newlist()
	for i := 0; i < 10; i++ {
		li.Insert(i)
		time.Sleep(1 * time.Second)
	}
	li.Traverse()
	fmt.Println("find 3, 6 ...")
	li.Find(3)
	time.Sleep(1 * time.Second)
	li.Find(6)
	li.Find(6)

	time.Sleep(1 * time.Second)
	li.Traverse()
	fmt.Println("Evict...")
	li.Find(0)
	li.Evict()
	li.Traverse()
}
