package main

import "fmt"

type Humaner interface {
	//方法，只有声明，没有实现,由别的类型（自定义类型）实现
	sayHi()
}

type Student struct {
	name string
	id   int
}

//Student实现了sayhi方法
func (s *Student) sayHi() {
	fmt.Printf("Student [%s , %d] sayhi\n", s.name, s.id)
}

type Teacher struct {
	addr  string
	group string
}

func (t *Teacher) sayHi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", t.addr, t.group)
}

type Mystr string

func (m *Mystr) sayHi() {
	fmt.Printf("Mystr[%s] sayhi\n", *m)
}

//定义一个普通函数，函数的参数是接口类型
//只有一个函数，却可以有不同的实现，这个叫多态
func WhoSayHi(i Humaner) {
	i.sayHi()
}

func main() {
	s := &Student{"mike", 666}
	t := &Teacher{"bj", "go"}
	var str Mystr = "hello mike"
	//调用同一个函数，不同表现，多态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, i := range x {
		i.sayHi()
	}
}
func main01() {
	//定义接口类型的变量
	var i Humaner
	//只要实现了此接口方法的类型，那么这个类型的变量（接收者类型）都可以给i赋值
	s := &Student{"mike", 666}
	i = s
	i.sayHi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayHi()

	var str Mystr = "hello mike"
	i = &str
	i.sayHi()
}
