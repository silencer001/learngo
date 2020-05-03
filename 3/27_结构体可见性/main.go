package main

import (
	"fmt"
	"structpackage"
)

func main() {
	s := structpackage.Student{Id: 1}
	s.Age = 10
	s.Sex = 'm'
	structpackage.Test(&s)
	fmt.Printf("%q", s)
}
