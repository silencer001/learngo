package main

import "fmt"

//实现2数相加
//面向过程
func Add01(a, b int) int {
	return a + b
}

//面向对象，方法：给某个类型绑定一个函数
type long int

func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = Add01(1, 2)
	fmt.Println("result:", result)
	var res2 long = 2
	a:=res2.Add02(12)
	fmt.Println(a)
}
