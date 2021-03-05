package btree

import (
	"fmt"

	"judeng.com/myqueue"
)

type Btree struct {
	root   *BtreeNode
	length int
}

type BtreeNode struct {
	payload int
	left    *BtreeNode
	right   *BtreeNode
}

/* PreOrder :前序遍历*/
func (btree *Btree) PreOrder() {
	preOrder(btree.root)
}

/*辅助函数*/
func preOrder(root *BtreeNode) {
	if root == nil { //终止条件
		return
	}
	fmt.Println(root.payload)
	preOrder(root.left)
	preOrder(root.right)
	return
}

/* InOrder :中序遍历*/
func (btree *Btree) InOrder() {
	inOrder(btree.root)
}
func inOrder(root *BtreeNode) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root.payload)
	inOrder(root.right)
	return
}

/*后序遍历*/

func (btree *Btree) PostOrder() {
	postOrder(btree.root)
}
func postOrder(root *BtreeNode) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Println(root.payload)
	return
}

/*BreadthFirstOrder : 按层广度优先遍历*/
func (btree *Btree) BreadthFirstOrder() {
	qu := myqueue.NewQueue()
	if btree.root == nil {
		return
	}
	qu.InQueue(btree.root)
	for !qu.IsEmpty() {
		node := qu.DeQueue().(BtreeNode)
		fmt.Println(node.payload)
		if node.left != nil {
			qu.InQueue(node.left)
		}
		if node.right != nil {
			qu.InQueue(node.right)
		}
	}
}
