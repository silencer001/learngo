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
	//超集可以转换为子集，反过来子集不能转换为超集
	var iPro Personer //超集
	iPro = &Student{"mike", 666}
	var i Humaner //子集
	// cannot use i (type Humaner) as type Personer in assignment:
	//    Humaner does not implement Personer (missing sing method)
	//iPro = i

	i = iPro
	i.sayHi()
}
