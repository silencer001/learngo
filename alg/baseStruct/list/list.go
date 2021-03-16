package list

type List struct {
	len  int
	head *ListNode
	tail *ListNode
}

type ListNode struct {
	data     int
	backward *ListNode
	forward  *ListNode
}

func NewList() *List {
	return &List{}
}

func (list *List) AddNodeTail(data int) {
	node := &ListNode{data: data}
	if list.tail != nil {
		list.tail.forward = node
	}
	node.backward = list.tail

	if list.head == nil {
		list.head = node
	}
	list.len++
}
func (list *List) DelNode(node *ListNode) {
	if node.backward != nil { //not head
		node.backward.forward = node.forward
	} else {
		list.head = node.forward
	}

	if node.forward != nil {
		node.forward.backward = node.backward
	} else {
		list.tail = node.backward
	}
	list.len--
}
