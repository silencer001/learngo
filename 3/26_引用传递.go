package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test(p *Student) {
	p.id = 888
	fmt.Println("test p", p)
}

func main() {
	s := Student{1, "mike", 'm', 18, "bj"}
	test(&s)
	fmt.Println("main s", s)
}
