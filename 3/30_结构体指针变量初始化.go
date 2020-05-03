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
	var p1 *Student = &Student{1, "ju", 'm', 18, "js"}
	fmt.Println("p1 = ", *p1)

	p2 := &Student{name: "mike", addr: "js"}
	fmt.Printf("p2 type:%T\n", p2)
	fmt.Println("p2 = ", p2)
	fmt.Printf("p2:%p", p2)
}
