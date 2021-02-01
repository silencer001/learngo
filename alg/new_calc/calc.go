package calc

import (
	"unicode"
)

//操作符优先级
var Operator map[rune]int = map[rune]int{
	'+': 1,
	'-': 1,
	'/': 2,
	'*': 2,
	'(': 0,
	')': 3,
}

type RpnExpress []interface{}

func Calc(s string) int {
	rpns := Reverse_polish_notation(s)
	return CalcRpn(rpns)
}

/*只支持数字、+、-、(、)、空格 */
/*首先将中序表达式转为逆波兰表达式
中缀表达式a + b*c + (d * e + f) * g，其转换成后缀表达式则为a b c * + d e * f  + g * +*/
func Reverse_polish_notation(s string) RpnExpress {
	rpnStack := NewStack()
	var rpn RpnExpress
	number := 0
	var appendflag bool = false
	for _, c := range s {
		if c == ' ' {
			continue
		}
		if unicode.IsDigit(c) {
			number = number*10 + (int)(c-48)
			appendflag = true
		} else {
			if appendflag == true {
				rpn = append(rpn, number)
				appendflag = false
				number = 0
			}
			processOperator(c, rpnStack, &rpn)
		}
	}
	//最后的参数
	if appendflag {
		rpn = append(rpn, number)
	}
	for {
		o, _ := rpnStack.Pop()
		if o == nil {
			break
		}
		rpn = append(rpn, o)
	}
	return rpn
}
func CalcParam(oper rune, param1 int, param2 int) int {
	switch oper {
	case '+':
		return param1 + param2
	case '-':
		return param1 - param2
	case '*':
		return param1 * param2
	case '/':
		return param1 / param2
	default:
		panic("unknow operator")
	}
}
func CalcRpn(rpn RpnExpress) int {
	rpnStack := NewStack()
	for _, x := range rpn {
		switch val := x.(type) {
		case int:
			rpnStack.Push(x)
		case rune:
			param2, _ := rpnStack.Pop()
			param1, e := rpnStack.Pop()
			if e != nil { //符号有可能是负数，假设没有非法参数
				param1 = 0
			}
			rpnStack.Push(CalcParam(val, param1.(int), param2.(int)))
		default:
			panic("unknown type")
		}
	}
	res, _ := rpnStack.Pop()
	return res.(int)
}

/*处理符号*/
func processOperator(c rune, rpn *Stack, rpns *RpnExpress) {
	if c == '(' {
		rpn.Push(c)
		return
	}
	//判断符号，当是)时，pop到rpn中
	if c == ')' {
		for {
			o, e := rpn.Pop()
			if e != nil {
				panic(e)
			}
			if o.(rune) != '(' {
				*rpns = append(*rpns, o)
			} else {
				return
			}
		}
	}
	//其他字符，当优先级大于栈顶元素或堆栈为空时push到堆栈中
	//当优先级小于等于栈顶元素时，弹出栈顶元素，并递归
	for {
		o, _ := rpn.Peek()
		if o == nil || Operator[c] > Operator[o.(rune)] {
			rpn.Push(c)
			break
		} else {
			*rpns = append(*rpns, o)
			rpn.Pop()
		}
	}
	return
}
