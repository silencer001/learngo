package main

import "fmt"

type mystr string
type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //结构体匿名字段
	int    //基础类型的匿名字段
	addr   string
	mystr
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 18, "bj", "test"}
	fmt.Printf("s = %=v\n", s)
	fmt.Println(s.name, s.int)
}
