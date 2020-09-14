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

func SearchBST(tree *BiTNode, key int, f *BiTNode) (*BiTNode, bool) {
	if tree == nil {
		return f, false
	} else if key == tree.data {
		return tree, true
	} else if key < tree.data {
		return SearchBST(tree.left, key, tree)
	} else {
		return SearchBST(tree.right, key, tree)
	}
}
func InsertBST(tree **BiTNode, key int) bool {
	p, ok := SearchBST(*tree, key, nil)
	if ok {
		fmt.Println("already insert value", key)
		return false //已经被插入
	}
	node := new(BiTNode)
	node.data = key
	if p == nil {
		//tree = &node，会导致tree最终没数据，为什么?
		*tree = node
		fmt.Println("debug:first insert value", (*tree).data)
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
	if (*tree).left == nil && (*tree).right == nil {
		fmt.Println("left and right is null, node:", (*tree).data)
		*tree = nil
	} else if (*tree).left == nil { //左子树为空
		*tree = (*tree).right
		fmt.Println("left is null")
	} else if (*tree).right == nil {
		*tree = (*tree).left
		fmt.Println("right is null")
	} else {
		prenode := (*tree).left
		prenode_father := *tree
		for prenode.right != nil { //查找前驱节点tmp
			prenode_father = prenode
			prenode = prenode.right
		}
		(*tree).data = prenode.data  //把前驱节点的值赋值给tree
		if prenode_father == *tree { //前驱节点就是左节点时，直接上移(此时没有右子树)
			(*tree).left = prenode.left
		} else { //前驱节点的左子树，变为前驱节点父节点的右子树
			prenode_father.right = prenode.left
		}
	}
}
func DeleteBST(tree **BiTNode, key int) bool {
	if tree == nil {
		return false
	}
	if key == (*tree).data {
		//fmt.Println("after delete:", (*tree).data)
		DeleteNode(tree)
		return true
	}
	if key < (*tree).data {
		return DeleteBST(&((*tree).left), key)
	} else {
		return DeleteBST(&((*tree).right), key)
	}
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
	var tree *BiTNode
	for _, val := range a {
		ok := InsertBST(&tree, val)
		fmt.Println("debug:insert", val, ok)
	}
	fmt.Println("中序遍历结果：")
	tree.OrderTree()
	delete, _ := strconv.Atoi(os.Args[1])
	DeleteBST(&tree, delete)
	fmt.Println("new")
	tree.OrderTree()
}
