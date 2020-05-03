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
	var s Student
	s.id = 1
	s.name = "mike"
	s.sex = 'm'
	s.addr = "bj"
	fmt.Println(s)

	var p *Student
	//指针必须有合法指向后，才能操作成员
	p = new(Student)
	p.id = 2
	p.name = "jude"
	p.sex = 'm'
	p.addr = "js"
	fmt.Println(*p)
}
