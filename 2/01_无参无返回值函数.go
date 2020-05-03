package main

import "fmt"

func fun() {
	fmt.Println("this is a function")
}

//定义了但没使用的函数，不会有告警
func fun1() {
	a := 66
	fmt.Println(a)
}
func main() {
	fun() //函数调用
}
