package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

/*模仿Redis的跳跃表实现*/

type SkipList struct {
	header *SkipListNode
	tail   *SkipListNode
	length int
	level  int
}
type SkipListNode struct {
	score    int64
	ele      string
	backward *SkipListNode
	level    []skipListLevel
}
type skipListLevel struct {
	forward *SkipListNode
	span    int //和forward的距离，注意，不是间隔节点个数
}

const (
	SKIPLIST_MAX_LEVEL int = 32
)

func NewSkipListNode(level int, score int64, ele string) *SkipListNode {
	node := new(SkipListNode)
	node.level = make([]skipListLevel, level) //每一层level不需要显式初始化
	node.score = score
	node.ele = ele
	return node
}
func NewSkipList() *SkipList {
	rand.Seed(time.Now().Unix())
	sl := new(SkipList)
	//还需要初始化header节点哦
	sl.header = NewSkipListNode(SKIPLIST_MAX_LEVEL, 0, "")
	sl.level = 1 //最少1层
	return sl
}

//辅助函数，返回1~32之间的整数
func randomLevel() int {
	level := 1
	for (rand.Int() % 0xFFFF) < int(0xFFFF>>2) { //1/4的概率
		level++
	}
	if level > SKIPLIST_MAX_LEVEL {
		level = SKIPLIST_MAX_LEVEL
	}
	return level
}

/*向跳跃表中插入元素*/
func (this *SkipList) SkipListInsert(score int64, ele string) {
	rank := make([]int, SKIPLIST_MAX_LEVEL)
	update := make([]*SkipListNode, SKIPLIST_MAX_LEVEL)
	/*首先从上到下找到待更新的每层节点及该节点的rank*/
	x := this.header
	for i := this.level - 1; i >= 0; i-- {
		if i == this.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil {
			if x.level[i].forward.score < score || (x.level[i].forward.score == score && x.level[i].forward.ele < ele) {
				rank[i] += x.level[i].span
				x = x.level[i].forward
			} else {
				break
			}
		}
		update[i] = x //rank[i]是第i层x的排名
		//fmt.Printf("debug:update[i]=%s\n", update[i].ele)
	}
	/*创建节点*/
	myLevel := randomLevel()
	fmt.Printf("mylevel:%d\n", myLevel)
	node := NewSkipListNode(myLevel, score, ele)
	/*更新update和rank*/
	if myLevel > this.level {
		for i := this.level; i < myLevel; i++ {
			update[i] = this.header
			rank[i] = 0
			update[i].level[i].span = this.length
		}
		this.level = myLevel
	}

	/*更新span和每层的forward指针*/
	for i := myLevel - 1; i >= 0; i-- {
		node.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = node

		//更新update[i]和node.level[i]的span
		node.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = rank[0] - rank[i] + 1
	}
	/*myLevel以上的层级，span需要加1*/
	for i := myLevel; i < this.level; i++ {
		update[i].level[i].span++
	}
	/*更新backward指针*/
	if update[0] == this.header {
		node.backward = nil
	} else {
		node.backward = update[0]
	}
	if node.level[0].forward != nil {
		node.level[0].forward.backward = node
	} else {
		this.tail = node
	}

	this.length++
}

/*
func (this *SkipList) SkipListFindNode(score int64, ele string) *SkipListNode {
	node := this.header
	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil {
			fmt.Printf("node.level[i].forward.score:%d, ele:%s\n", node.level[i].forward.score, node.level[i].forward.ele)
			if node.level[i].forward.score < score || (node.level[i].forward.score == score && node.level[i].forward.ele < ele) {
				node = node.level[i].forward
			} else if node.level[i].forward.score == score && node.level[i].forward.ele == ele {
				return node
			} else {
				break
			}
		}
	}
	return nil
}*/
func (this *SkipList) SkipListFindNode(score int64, ele string) *SkipListNode {
	node := this.header
	for i := this.level - 1; i >= 0; i-- {
		for node.level[i].forward != nil {
			if node.level[i].forward.score < score || (node.level[i].forward.score == score && node.level[i].forward.ele < ele) {
				node = node.level[i].forward
			} else {
				break
			}
		}
	}
	node = node.level[0].forward
	if node != nil && node.score == score && node.ele == ele {
		return node
	} else {
		return nil
	}
}

func (this *SkipList) DebugPrintSkipList() {
	node := this.header
	for node != nil {
		fmt.Printf("level:%d,score=%d,ele=%s\n", len(node.level), node.score, node.ele)
		for i := len(node.level) - 1; i >= 0; i-- {
			var forward string
			if node.level[i].forward == nil {
				forward = "nil"
			} else {
				forward = node.level[i].forward.ele
			}
			fmt.Printf(" level[%d]:span=%d,forward=%v |", i, node.level[i].span, forward)
		}
		fmt.Printf("\n")
		node = node.level[0].forward
	}
}
