package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (person Person) SetInfoValue() {
	fmt.Println("SetInfoValue", person)
}
func (person *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer", *person)
}

func main() {
	//假如结构体是一个指针变量，它能够调用哪些方法
	//这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 18}

	p.SetInfoPointer()
	//内部做了转行，把指针p，转成(*p)
	//等同于(*p).SetInfoValue
	p.SetInfoValue()

	v := Person{"yoyo", 'f', 20}
	//内部做了转换，把值v，转换成(&v)
	v.SetInfoPointer()
	(&v).SetInfoPointer()
	v.SetInfoValue()

}
