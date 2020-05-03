package main

import "fmt"

type mystr string
type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*Person     //结构体匿名字段
	id      int //基础类型的匿名字段
	addr    string
}

func main() {
	s1 := Student{&Person{"Mike", 'm', 18}, 666, "js"}
	//s1 =  {0xc0000044c0 666 js}
	fmt.Println("s1 = ", s1)

	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

	var s2 Student
	s2.Person = new(Person) //分配空间
	s2.name = "yoyo"
	s2.sex = 'm'
	s2.age = 20
	s2.id = 12
	s2.addr = "bj"
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.addr)
}
