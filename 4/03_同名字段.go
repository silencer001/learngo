package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
	name string
}

func main() {
	var s Student
	//默认规则(就近原则)
	/*如果本作用域找到此成员，就操作此成员
	  如果没有找到，找到继承的字段* /
	/*同层次，不允许有相同的成员名称*/
	s.name = "Mike"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	s.Person.name = "test"
	fmt.Printf("s.name:%s", s.name)
}
