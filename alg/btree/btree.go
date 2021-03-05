package btree

import (
	"fmt"
	"github.com/silencer001/learngo/alg/btree/internal/myqueue"
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

/*前序遍历*/
func (this *Btree) PreOrder() {
	preOrder(this.root)
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

/*中序遍历*/
func (this *Btree) InOrder() {
	inOrder(this.root)
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

func (this *Btree) PostOrder() {
	postOrder(this.root)
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

/*按层广度优先遍历*/
func (this *Btree) BreadthFirstOrder() {
	qu := myqueue.NewQueue()
}
