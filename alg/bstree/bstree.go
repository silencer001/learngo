package bstree

/*二叉搜索树的相关api*/
import "judeng.com/btree"

type Bstree struct {
	btree.Btree
}

func (bst *Bstree) Search(v int) *btree.BtreeNode {
	node := bst.Root

	for node != nil {
		if node.Payload == v {
			return node
		} else if node.Payload > v {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}

/*非递归方式向二叉搜索树上插入节点*/
func (bst *Bstree) Instert(v int) {
	nnode := btree.NewBtreeNode(v)
	node := bst.Root
	for node != nil {
		if node.Payload == v {
			return
		} else if node.Payload > v {
			if node.Left == nil {
				node.Left = nnode
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = nnode
				break
			}
			node = node.Right
		}
	}
	bst.Length++
	return
}

/*非递归方式删除二叉搜索树上的节点*/
func (bst *Bstree) Delete(v int) {
	/*首先查找该节点及其父节点*/
	var isleft bool
	var pnode *btree.BtreeNode = nil
	node := bst.Root
	for node != nil {
		if node.Payload == v {
			break
		} else if node.Payload > v {
			pnode = node
			isleft = true
			node = node.Left
		} else {
			pnode = node
			isleft = false
			node = node.Left
		}
	}
	//not found
	if node == nil {
		return
	}
	delete := func(pnode, snode *btree.BtreeNode, isleft bool) {
		if pnode == nil {
			bst.Root = snode
		} else {
			if isleft {
				pnode.Left = snode
			} else {
				pnode.Right = snode
			}
		}
		bst.Length--
	}
	/*分为三种情况删除该节点*/
	if node.Left == nil && node.Right == nil {
		delete(pnode, nil, isleft)
	} else if node.Left != nil && node.Right == nil {
		delete(pnode, node.Left, isleft)
	} else if node.Left == nil && node.Right != nil {
		delete(pnode, node.Right, isleft)
	} else { //待删除节点有左右子树时，需要找到右子树上的最小值
		findmin := func(root *btree.BtreeNode) (min, minp *btree.BtreeNode) {
			min = root
			minp = nil
			for min.Left != nil {
				minp = min
				min = min.Left
			}
			return min, minp
		}
		min, minp := findmin(node.Right)
		if min.Right == nil { //min必定没有左子树，此时只需用判断是否有右子树
			if minp != nil {
				minp.Left = nil
			}
		} else {
			if minp != nil {
				minp.Left = min.Right
			}
		}
		min.Right = node.Right
		min.Left = node.Right
		pnode.Right = min
		bst.Length--
	}
}
