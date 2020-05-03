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
	s := Student{}
	s.name = "jude"
	s.sex = 'm'
	s.Person.age = 18
	s.id = 66
	fmt.Println(s)
	//顺序初始化
	s2 := Student{
		Person{"mike", 'm', 19},
		888,
		"bj",
	}
	fmt.Println(s2)

	s3 := Student{
		Person: Person{"mike", 'm', 19},
		addr:   "bj",
	}
	fmt.Printf("%+v", s3)
}
