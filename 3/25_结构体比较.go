package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	s1 := Student{1, "Mike", 'm', 18, "bj"}
	s2 := Student{1, "Mike", 'm', 18, "bj"}
	s3 := Student{2, "Mike", 'm', 18, "bj"}
	fmt.Println("s1 == s2", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)

	var s4 Student
	s4 = s3
	s4.age = 19
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
}
