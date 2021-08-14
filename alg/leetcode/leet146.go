package leetcode

import "container/list"

type LRUCache struct {
	cap     int
	dict    map[int]*list.Element
	lruList *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:     capacity,
		dict:    map[int]*list.Element{},
		lruList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.dict[key]
	if !ok {
		return -1
	}
	//e.Prev().Next()= e.Next()
	this.lruList.MoveToFront(e)
	return e.Value.(int)
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.dict[key]
	if ok {
		e.Value = value
		this.lruList.MoveToFront(e)
		return
	}
	if this.lruList.Len() == this.cap {
		this.lruList.Remove(this.lruList.Back())
	}
	e = this.lruList.PushFront(value)
	this.dict[key] = e
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
