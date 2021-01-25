package mystack

import (
	"errors"
	"fmt"
)

type Stack []interface{}

/* return a Satck type，with len n*/
func NewStack(n int) *Stack {
	s := make([]interface{}, 0, n)
	p := &s
	return (*Stack)(p)
}

func (s *Stack) PushStack(val interface{}) error {
	if len(*s) >= cap(*s) {
		return errors.New("stack is full")
	}
	fmt.Println("len(s):", len(*s), "cap(s)=", cap(*s))
	*s = append(*s, val)
	return nil
}

func (s *Stack) PopStack() (interface{}, error) {
	if len(*s) <= 0 {
		return nil, errors.New("s is empty")
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, nil
}

func (s *Stack) IsEmpty() bool {
	if len(*s) > 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) TopStack() (interface{}, error) {
	if len(*s) <= 0 {
		return nil, errors.New("s is empty")
	}
	return (*s)[len(*s)-1], nil
}
