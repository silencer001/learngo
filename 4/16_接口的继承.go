package main

import "fmt"

type Humaner interface { //子集
	//方法，只有声明，没有实现,由别的类型（自定义类型）实现
	sayHi()
}

type Personer interface {
	Humaner //匿名字段
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

//Student实现了sayhi方法
func (s *Student) sayHi() {
	fmt.Printf("Student [%s , %d] sayhi\n", s.name, s.id)
}

func (s *Student) sing(lrc string) {
	fmt.Println("Student在唱歌: ", lrc)
}

func main() {
	var i Personer
	s := &Student{"mike", 666}
	i = s
	i.sayHi()
	i.sing("rolling in the deep")
}
