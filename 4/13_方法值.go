package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (person Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &person, person)
}
func (person *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", person, *person)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main :%p, %v\n", &p, p)

	p.SetInfoPointer()
	//保存方法入口地址
	pFunc := p.SetInfoPointer //这个就是方法值，调用函数时无需再传递接收者
	pFunc()
	fmt.Printf("pfunc type:%T\n", pFunc)

	vFunc := p.SetInfoValue
	vFunc()
	fmt.Printf("vfunc type:%T\n", vFunc)
}
