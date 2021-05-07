package btree

import (
	"fmt"

	"judeng.com/myqueue"
	"judeng.com/mystack"
)

type Btree struct {
	Root   *BtreeNode
	Length int
}

type BtreeNode struct {
	Payload int
	Left    *BtreeNode
	Right   *BtreeNode
}

func NewBtreeNode(v int) *BtreeNode {
	return &BtreeNode{
		Payload: v,
	}
}

/* PreOrder :前序遍历*/
func (btree *Btree) PreOrder() {
	preOrder(btree.Root)
}

/*辅助函数*/
func preOrder(root *BtreeNode) {
	if root == nil { //终止条件
		return
	}
	fmt.Println(root.Payload)
	preOrder(root.Left)
	preOrder(root.Right)
	return
}

/* InOrder :中序遍历*/
func (btree *Btree) InOrder() {
	inOrder(btree.Root)
}
func inOrder(root *BtreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Println(root.Payload)
	inOrder(root.Right)
	return
}

/*非递归方式中序遍历*/
func (btree *Btree) InOrder() {
	p := btree.Root
	s := mystack.NewStack()
	for p != nil || !s.IsEmpty() {

	}
}

/*后序遍历*/

func (btree *Btree) PostOrder() {
	postOrder(btree.Root)
}
func postOrder(root *BtreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Println(root.Payload)
	return
}

/*BreadthFirstOrder : 按层广度优先遍历*/
func (btree *Btree) BreadthFirstOrder() {
	qu := myqueue.NewQueue()
	if btree.Root == nil {
		return
	}
	qu.InQueue(btree.Root)
	for !qu.IsEmpty() {
		node := qu.DeQueue().(BtreeNode)
		fmt.Println(node.Payload)
		if node.Left != nil {
			qu.InQueue(node.Left)
		}
		if node.Right != nil {
			qu.InQueue(node.Right)
		}
	}
}
