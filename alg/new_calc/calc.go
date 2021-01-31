package calc

import (
	"unicode"
)
//操作符优先级
const Operator map[byte] int {
	'+':1
	'-':1
	'/':2
	'*':2
	'(':3
}
type RpnExpress []interface{}

/*只支持数字、+、-、(、)、空格 */
func Reverse_polish_notation(s string) RpnExpress {
	reverse_polish_notation := NewStack()
	var rpn RpnExpress
	number := 0
	appendflag = false
	for _, c:= range s {
		if c == ' ' {
			continue
		}
		if unicode.IsDigit(c) {
			number = number*10 + (int)(c-48)
			appendflag = true
		} else {
			if appendflag == true  {
				rpn = append(rpn, number)
				appendflag = false
			}
			//push到符号栈中
			
		}
	}
}
