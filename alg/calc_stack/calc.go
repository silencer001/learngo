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
