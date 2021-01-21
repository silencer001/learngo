package main

import (
	"fmt"
	"os"
	"strconv"
)

type BiTNode struct {
	data  int
	right *BiTNode
	left  *BiTNode
}
type BiTree struct {
	root *BiTNode
	n    int
}

/*hint:return node and true
  miss:return nil and false
*/
func SearchBST(tree *BiTNode, key int, parent *BiTNode) (*BiTNode, *BiTNode, bool) {
	if tree == nil {
		return nil, parent, false
	} else if key == tree.data {
		return tree, parent, true
	} else if key < tree.data {
		return SearchBST(tree.left, key, tree)
	} else {
		return SearchBST(tree.right, key, tree)
	}
}
func (this *BiTree) InsertBST(key int) bool {
	_, p, ok := SearchBST(this.root, key, nil)
	if ok {
		fmt.Println("already insert value", key)
		return false //已经被插入
	}
	node := new(BiTNode)
	node.data = key
	if p == nil {
		this.root = node
		this.n++
		fmt.Println("debug:first insert value", this.root.data)
		return true
	}
	if p.data > key {
		p.left = node
	} else {
		p.right = node
	}
	return true

}

func DeleteNode(tree **BiTNode) {
	node := *tree
	if node.left == nil && node.right == nil {
		fmt.Println("left and right is null, node:", node.data)
		*tree = nil
	} else if node.left == nil { //左子树为空
		*tree = node.right
		fmt.Println("left is null")
	} else if node.right == nil {
		*tree = node.left
		fmt.Println("right is null")
	} else {
		prenode := node.left
		prenodeFather := node
		for prenode.right != nil { //查找前驱节点tmp
			prenodeFather = prenode
			prenode = prenode.right
		}
		node.data = prenode.data   //把前驱节点的值赋值给tree
		if prenodeFather == node { //前驱节点就是左节点时，直接上移(此时前驱节点没有右子树)
			node.left = prenode.left
		} else { //前驱节点的左子树，变为前驱节点父节点的右子树
			prenodeFather.right = prenode.left
		}
	}
}

func (this *BiTree) DeleteBST(key int) bool {
	node, parent, ok := SearchBST(this.root, key, nil)
	if !ok {
		return false
	}
	if parent == nil {
		DeleteNode(&this.root)
	} else {
		if parent.left == node {
			DeleteNode(&parent.left)
		} else {
			DeleteNode(&parent.right)
		}
	}
	this.n--
	return true
}
func (t *BiTNode) OrderTree() {
	if t == nil {
		return
	}
	t.left.OrderTree()
	fmt.Println(t.data)
	t.right.OrderTree()
}

func main() {
	a := []int{100, 3, 4, 2, 5, 9, 1, 6, 8, 12, 13, 10}
	tree := new(BiTree)
	for _, val := range a {
		ok := tree.InsertBST(val)
		fmt.Println("debug:insert", val, ok)
	}
	fmt.Println("中序遍历结果：")
	tree.root.OrderTree()
	delete, _ := strconv.Atoi(os.Args[1])
	tree.DeleteBST(delete)
	fmt.Println("new")
	tree.root.OrderTree()
}
