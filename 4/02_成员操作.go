package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s1 := Student{Person{"MIKE", 'm', 18}, 123, "bj"}
	s1.name = "jude"
	s1.age += 1
	s1.Person = Person{"test", 'a', 20}
	fmt.Printf("s1 %+v\n", s1)
}
