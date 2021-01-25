package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack(3)
	s.PushStack(1)
	s.PushStack(2)
	s.PushStack(3)
	s.PushStack(4)
	fmt.Println(s.PopStack())
	fmt.Println(s.PopStack())
	fmt.Println(s.PopStack())
	fmt.Println(s.PopStack())

}
