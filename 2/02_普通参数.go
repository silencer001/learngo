package main

import "fmt"

//有参无返回值函数的定义，普通参数列表
//定义函数时，在函数名后面()定义的参数叫形参
//参数传递，只能由实参传递给形参，不能反向传递
func Myfunc01(i int, j int) {
	fmt.Println("i+j=", i+j)
}
func Myfunc02(i, j int) {
	fmt.Println("i-j=", i-j)
}

func Myfunc04(i int, j float64) {
	fmt.Println("i+j=", i-int(j))
}
func main() {
	Myfunc01(1, 2)
}
