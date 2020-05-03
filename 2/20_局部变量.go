package main

import "fmt"

func test() {
	a := 10
	fmt.Println("a=", a)
}
func main() {
	//定义在{}里边的变量就是局部变量，只能在{}里有效
	//作用域，变量其作用的范围
	{
		i := 10
	}
	//undefined: i
	//fmt.Println("i=", i)
	//变量的作用域是语句块，if for {}
	if flag :=3, flag == 3{
		fmt.Println("flag", flag)
	}
	//undefined:flag
	//flag = 4
}
