package calc

import (
	"errors"
	"mystack"
	"strconv"
	"strings"
)

var prio_operator map[byte]int = map[byte]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

func PushPram(str string, stack *mystack.Stack) (string, error) {
	if len(str) == 0 {
		return "", errors.New("input string is empty")
	}
	var number int
	var e error
	pramIndex := strings.IndexAny(str, "+-*/")
	if pramIndex == -1 {
		number, e = strconv.Atoi(str)
	} else {
		number, e = strconv.Atoi(str[:pramIndex])
	}
	if e != nil {
		panic(e)
	}
	stack.PushStack(number)
	if pramIndex == -1 {
		return "", nil
	}
	return str[pramIndex:], nil
}
func InputExpress(s string) int {
	prams := mystack.NewStack(100)
	opras := mystack.NewStack(100)
	var e error
	s, e = PushPram(s, prams)
	if e != nil {
		panic(e)
	}
	for s != "" {
		switch byte(s[0]) {
		case '+':
			fallthrough
		case '-':
			fallthrough
		case '*':
			fallthrough
		case '/':
			operatop, _ := opras.TopStack()
			if operatop == nil || prio_operator[s[0]] < prio_operator[operatop.(byte)] {
				opras.PushStack(s[0])
				s, e = PushPram(s[1:], prams)
			} else {
				var op byte = s[0]
				s, e = PushPram(s[1:], prams)
				if e != nil {
					panic(e)
				}
				param2, _ := prams.PopStack()
				param1, _ := prams.PopStack()
				res := MustCalc(param1.(int), param2.(int), op)
				prams.PushStack(res)
			}
		default:
			panic("unknow or unsupport operator!")
		}
	}
	var result int
	for {
		opera, _ := opras.PopStack()
		if opera == nil {
			break
		}
		param2, _ := prams.PopStack()
		param1, _ := prams.PopStack()
		result = MustCalc(param1.(int), param2.(int), opera.(byte))
		prams.PushStack(result)
	}
	return result
}

func MustCalc(param1, param2 int, opera byte) int {
	switch opera {
	case '+':
		return param1 + param2
	case '-':
		return param1 - param2
	case '*':
		return param1 * param2
	case '/':
		return param1 / param2
	default:
		panic("unkown operator")
	}
}

type Stack []int

func (this *Stack) Push(x int) {
	*this = append(*this, x)
}
func (this *Stack) Pop() int {
	if len(*this) == 0 {
		return 0
	}
	res := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return res
}
func (this *Stack) Peek() int {
	if len(*this) == 0 {
		return 0
	}
	return (*this)[len(*this)-1]
}
func (this *Stack) IsEmpty() bool {
	if len(*this) == 0 {
		return true
	} else {
		return false
	}
}

type MyQueue struct {
	istack Stack
	ostack Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		istack: []int{},
		ostack: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.istack.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.ostack.IsEmpty() {
		for !this.istack.IsEmpty() {
			x := this.istack.Pop()
			this.ostack.Push(x)
		}
	}
	return this.ostack.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.ostack.IsEmpty() {
		for !this.istack.IsEmpty() {
			x := this.istack.Pop()
			this.ostack.Push(x)
		}
	}
	return this.ostack.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.ostack.IsEmpty() && this.istack.IsEmpty() {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
