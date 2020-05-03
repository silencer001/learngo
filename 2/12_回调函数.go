package main

import "fmt"

type functype func(int, int) int

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
//计算器，进行四则运算

func calc(a, b int, fTest functype) (result int) {
	fmt.Println("Calc ")
	result = fTest(a, b)
	return
}

func main() {

}
