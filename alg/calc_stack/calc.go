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
	for s != "" {
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
			operatop, _ := opras.TopStack()
			if prio_operator[s[0]] < prio_operator[operatop.(byte)] {
				opras.PushStack(s[0])
				s = s[1:]
			} else {
				s, e = PushPram(s[1:], prams)
				if e != nil {
					panic(e)
				}
				param2, _ := prams.PopStack()
				param1, _ := prams.PopStack()
				res := MustCalc(param1.(int), param2.(int), s[0])
				prams.PushStack(res)

			}
		default:
			panic("unknow or unsupport operator!")
		}
	}
	return 0 //dummy for compile
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
