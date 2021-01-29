package calc

import (
	"errors"
)

type Stack []interface{}

func NewStack() *Stack {
	return &Stack{}
}

func (this *Stack) IsEmpty() bool {
	if len(*this) > 0 {
		return false
	} else {
		return true
	}
}

func (this *Stack) Push(x interface{}) {
	*this = append(*this, x)
	return
}

func (this *Stack) Pop() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	ret := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return ret,nil
}

func (this *Stack) Peek() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return (*this)[len(*this)-1],nil
}

