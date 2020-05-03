package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (tmp *Person) PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d\n", tmp.name, tmp.sex, tmp.age)
}

//有一个学生，继承Person字段，成员和方法都继承
type Student struct {
	Person
	id   int
	addr string
}

func (tmp *Student) PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d, id = %d, addr = %s\n",
		tmp.name, tmp.sex, tmp.age, tmp.id, tmp.addr)
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "js"}
	s.PrintInfo()
}
