package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//带有接收者的函数叫方法

func (tmp *Person) PrintInfo() {
	fmt.Println("tmp = ", *tmp)
}

func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}
func main() {
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	p2 := new(Person)
	p2.SetInfo("yoyo", 'a', 20)
	p2.PrintInfo()

	var p3 Person
	p3.SetInfo("test", 'f', 21)
	p3.PrintInfo()
}
