package mystack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
