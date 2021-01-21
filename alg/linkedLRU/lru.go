package lru

import (
	"fmt"
	"time"
)

type Node struct {
	payload int
	lru     time.Time
	prev    *Node
	next    *Node
}
type List struct {
	len  int
	head *Node
	tail *Node
}

func (li *List) Find(value int) {
	for iter := li.head; iter != nil; iter = iter.next {
		if iter.payload == value {
			li.UpdateLRU(iter)
			break
		}
	}
	//do nothing
}
func (li *List) UpdateLRU(now *Node) {
	now.lru = time.Now()
	if now == li.head {
		return //已经是在链表头了
	}

	if now == li.tail {
		li.tail = now.prev
		li.tail.next = nil
	} else {
		//fmt.Printf("now.prev.next:%v\n", now.prev)
		//fmt.Printf("now.next:%v\n", now.next)
		now.prev.next = now.next //将prev的next指向now的next
		now.next.prev = now.prev
	}

	now.next = li.head
	li.head = now
}

func (li *List) Insert(payload int) {
	nd := new(Node)
	nd.payload = payload
	nd.lru = time.Now()
	if li.head != nil {
		li.head.prev = nd
	}
	nd.next = li.head
	nd.prev = nil
	li.head = nd
	if li.tail == nil {
		li.tail = nd
	}
	li.len++
}

func (li *List) Evict() {
	if li.len == 0 {
		return
	}
	if li.len == 1 {
		li.tail = nil
		li.head = nil
		li.len--
		return
	}
	li.tail = li.tail.prev
	li.tail.next = nil
	li.len--
}

func (li *List) Traverse() {
	for iter := li.head; iter != nil; iter = iter.next {
		fmt.Println(iter.payload, "lru clock: ", iter.lru.String())
	}
}
func Newlist() *List {
	li := new(List)
	li.len = 0
	li.head = nil
	li.tail = nil
	return li
}
