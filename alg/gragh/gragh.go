package gragh

import (
	"fmt"

	"github.com/silencer001/learngo/alg/baseStruct/list"
	"github.com/silencer001/learngo/alg/baseStruct/queue"
)

type Gragh struct {
	v    int
	edge []list.List
}

func NewGragh(spot int) *Gragh {
	g := &Gragh{v: spot}
	g.edge = make([]list.List, spot)
	return g
}
func (gragh *Gragh) AddEdge(s, t int) {
	gragh.edge[s].AddNodeTail(t)
	gragh.edge[t].AddNodeTail(s)
}

// Bfs : 广度遍历
func (gragh *Gragh) Bfs(s, t int) []int {
	if s == t {
		return []int{s}
	}
	visited := make([]bool, gragh.v) //点是否被访问过
	//该节点的上一个节点
	prev := make([]int, gragh.v)
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}
	que := queue.NewQueue()
	que.InQueue(s)
	visited[s] = true
	var curSpot int = -1
	//队列不为空时
	for !que.IsEmpty() {
		q := que.DeQueue().(int)
		for iters := gragh.edge[q].NewIter(); iters != nil; iters = gragh.edge[q].IterNext(iters) {
			curSpot = iters.Data()
			if visited[curSpot] {
				continue
			}
			visited[curSpot] = true
			prev[curSpot] = q
			if curSpot == t {
				break //found
			}
			que.InQueue(curSpot)
		}
		if curSpot == t {
			break
		}
	}
	if curSpot != t {
		return nil
	}
	return createResult(prev, curSpot)
}

func createResult(prev []int, t int) []int {
	var res []int
	res = append(res, t)
	for prev[t] != -1 {
		res = append(res, prev[t])
		t = prev[t]
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

// Dfs : 深度优先遍历
func (gragh *Gragh) Dfs(s, t int) []int {
	visited := make([]bool, gragh.v)
	prev := make([]int, gragh.v)
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}
	visited[s] = true
	type functype func(s, t int)
	var dfs functype //必须先声明函数变量，再递归调用，否则编译不通过
	found := false
	dfs = func(s, t int) {
		if found {
			return
		}
		curlist := gragh.edge[s]
		for iter := curlist.NewIter(); iter != nil; iter = curlist.IterNext(iter) {
			v := iter.Data()
			if visited[v] {
				continue
			}
			visited[v] = true
			prev[v] = s
			if v == t {
				found = true
				return
			}
			dfs(v, t)
		}
		return
	}
	dfs(s, t)
	if !found {
		return nil
	}

	fmt.Println(prev)
	return createResult(prev, t)
}
