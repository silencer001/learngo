package btree

import (
	"fmt"
	"testing"
)

func TestInOrder(t *testing.T) {
	bt := NewBtree()
	bt.Root = NewBtreeNode(5)
	bt.Root.Left = NewBtreeNode(3)
	bt.Root.Right = NewBtreeNode(7)
	bt.Root.Left.Right = NewBtreeNode(4)
	bt.Root.Right.Left = NewBtreeNode(6)
	bt.Root.Right.Right = NewBtreeNode(8)
	bt.InOrder()
	fmt.Println()
	bt.InOrderNorecur()
	fmt.Println()
	bt.PreOrder()
	fmt.Println()
	bt.PreOrderNorecur()
}
