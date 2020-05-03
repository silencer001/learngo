package main

import "fmt"
func main() {
	var a int 

	//每个变量有2层含义：变量的内存，变量的地址
	fmt.Printf("a = %d\n", a)
	fmt.Printf("P(a) = %p\n", &a)
	fmt.Printf("P(a) = %v\n", &a)

	//保存某个变量的地址，需要指针类型
	b := &a
	fmt.Printf("b = %v, type:%T\n", b,b)
	fmt.Printf("*b = %d\n", *b)

	*b = 10
	fmt.Printf("*b = %d\n", *b)
	fmt.Printf("a = %d\n", a)


	//不能操作没有内存的指针
	var p *int
	fmt.Printf("p = %v\n", p)
	//panic，p指向nil，没有内存空间不能赋值
	*p = 666
}