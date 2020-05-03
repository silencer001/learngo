package main

import "fmt"

func main() {
	a := 10
	str := "mike"
	//匿名函数，没有函数名字
	//这个函数可以捕获函数外面的变量
	f1 := func() { //自动推导函数类型
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	fmt.Println("main")
	f1()

	//
	type funcType func() //函数没有参数，没有返回值
	//声明函数变量
	var f2 funcType
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func() {
		a--
		fmt.Printf("a=%d\n", a)
	}() //记得加括号，func literal evaluated but not used
	fmt.Printf("in main a=%d\n", a)
	//带参数的匿名函数
	func(a int) {
		a--
		fmt.Printf("a = %d\n", a)
	}(a)
	fmt.Printf("in main a = %d\n", a) //仍然是9
	//有参有返回值的匿名函数
	i, j := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(20, 10)
	fmt.Printf("i = %d, j= %d\n", i, j)
}
