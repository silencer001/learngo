package main

import "fmt"

type Student struct {
	id   int //结构体成员声明时不能加var关键字
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//顺序初始化，每个成员都必须初始化
	var s1 Student = Student{
		1,
		"mike",
		'm',
		18,
		"js",
	}
	//指定成员初始化，没有指定的成员，自动赋值为默认值
	fmt.Println("s1 = ", s1)
	var s2 Student = Student{
		name: "judeng",
	}
	fmt.Println("s2 = ", s2)
}
