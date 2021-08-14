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

func NewBtree() *Btree {
	return &Btree{
		Root:   nil,
		Length: 0,
	}
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

/*非递归方式前序遍历*/
func (btree *Btree) PreOrderNorecur() {
	p := btree.Root
	s := mystack.NewStack(1024)
	for p != nil || !s.IsEmpty() {
		for p != nil {
			fmt.Println(p.Payload)
			s.Push(p)
			p = p.Left
		}
		if !s.IsEmpty() {
			x, _ := s.Pop()
			p = x.(*BtreeNode)
			p = p.Right
		}
	}
}

/*非递归方式中序遍历*/
func (btree *Btree) InOrderNorecur() {
	p := btree.Root
	s := mystack.NewStack(1024)
	for p != nil || !s.IsEmpty() {
		for p != nil {
			s.Push(p)
			p = p.Left
		}
		if !s.IsEmpty() {
			x, _ := s.Pop()
			p = x.(*BtreeNode)
			fmt.Println(p.Payload)
			p = p.Right
		}
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

/*先写个后续遍历的循环方法吧*/

func PostOrder2(root *BtreeNode) {
	fmt.Println()
	stack := make([]*BtreeNode, 0)
	p := root
	var prev *BtreeNode
	for p != nil || len(stack) > 0 {
		for p != nil {
			fmt.Println("Push: ", p.Payload)
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println("Pop: ", p.Payload)
			if (p.Right == nil) || p.Right == prev { //已经遍历过p
				fmt.Printf("visited:%d\n", p.Payload)
				prev = p
				p = nil
			} else {
				fmt.Println("Push: ", p.Payload)
				stack = append(stack, p) //重新push后，遍历右子树
				p = p.Right
			}
		}
	}
	fmt.Println()
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

/*DepthFirstOrder : 深度优先遍历 */

func (btree *Btree) DepthFirstOrder() {
	s := mystack.NewStack(100)
	s.Push(btree.Root)
	for !s.IsEmpty() {
		x, _ := s.Pop()
		p := x.(*BtreeNode)
		fmt.Println(p.Payload)
		if p.Left != nil {
			s.Push(p.Left)
		}
		if p.Right != nil {
			s.Push(p.Right)
		}
	}
}
