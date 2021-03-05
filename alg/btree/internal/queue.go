package myqueue

/*Queue : 内部队列结构*/
type Queue struct {
	len   int
	nodes []interface{}
}

/*NewQueue :构建空队列*/
func NewQueue() *Queue {
	return &Queue{
		len:   0,
		nodes: []interface{}{},
	}
}

/*InQueue : 入队函数，不考虑异常*/
func (queue *Queue) InQueue(node interface{}) {
	queue.nodes = append(queue.nodes, node)
	queue.len++
}

/*DeQueue :出队函数，不考虑异常*/
func (queue *Queue) DeQueue() interface{} {
	x := queue.nodes[0]
	queue.nodes = queue.nodes[1:]
	queue.len--
	return x
}

/*IsEmpty :判断队列是否为空*/
func (queue *Queue) IsEmpty() bool {
	if queue.len == 0 {
		return true
	} else {
		return false
	}
}
