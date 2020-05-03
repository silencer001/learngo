package main

import "fmt"

func test01() int {
	var x int
	x++
	return x * x //函数调用完毕，x被释放
}

//函数的返回值是匿名函数，返回一个函数类型
//变量的声明周期不是由作用域决定
func test02() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func test() {
	//test02函数中的x的变量生命周期是test函数
	f:=test02()
	fmt.Println("test func:",f())
	fmt.Println("test func:",f())
}
func main() {
	fmt.Println("first test")
	test();

	//返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	//只要闭包还在使用中，变量就存在
	//在这个实例中，闭包函数赋值给了f函数变量，所以一直在使用中
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("second test")
	test();
}

// func main() {
// 	fmt.Println(test01())
// 	fmt.Println(test01())
// 	fmt.Println(test01())
// 	fmt.Println(test01())
// }
