package list

import (
	"fmt"
	"testing"
)

func debugPrintlist(l *List) {
	iter := l.head
	for iter != nil {
		fmt.Println(iter.data)
		iter = iter.forward
	}
}
func TestAddNodeTail(t *testing.T) {
	l := NewList()
	l.AddNodeTail(1)
	l.AddNodeTail(2)
	l.AddNodeTail(3)
	l.AddNodeTail(4)
	debugPrintlist(l)
	n := l.FindNode(4)
	l.DelNode(n)
	n = l.FindNode(1)
	l.DelNode(n)
	debugPrintlist(l)
}
