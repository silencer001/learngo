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

	//方法值 f:= p.setInfoPointer //隐藏了接收者
	//方法表达式
	f := (*Person).SetInfoPointer
	f(&p)
	fmt.Printf("f type :%T\n", f)
	f2 := (Person).SetInfoValue
	//f(p) cannot use p (type Person) as type *Person in argument to f
	f2(p)
	fmt.Printf("f2 type :%T\n", f2)
}
