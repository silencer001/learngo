package calc

import (
	"errors"
	"mystack"
	"strconv"
	"strings"
)
var prio_operator map[byte]int = {
	'+':1,
	'-':1,
	'*':2,
	'/':2,
}
func PushPram(str string, stack *mystack.Stack) (string, error) {
	if len(str) == 0 {
		return "", errors.New("input string is empty")
	}
	pramIndex := strings.IndexAny(str, "+-*/")
	number, e := strconv.Atoi(str[:pramIndex-1])
	if e != nil {
		return "", e
	}
	stack.PushStack(number)
	return str[pramIndex:], nil
}
func InputExpress(s string) int {
	prams := mystack.NewStack(100)
	opras := mystack.NewStack(100)
	for s != nil {
		s, e := PushPram(s, prams)
		if e != nil {
			panic(e)
		}
		//最后一个参数
		if s == "" {
			break
		}
		switch s[0] {
		case '+':
		case '-':
		case '*':
		case '/':
		if prio_operator[s[0]] < prio_operator[opras.TopStack()] {
			opras.PushStack(s[0])
			s = s[1:]
		} else {
			s,e = PushPram(s[1:], prams)
			if e != nil {
				panic(e)
			}
		}
		default:
			panic("unknow or unsupport operator!")
		}
	}
}
