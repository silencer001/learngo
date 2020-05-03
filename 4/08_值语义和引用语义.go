package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (person Person) SetInfoValue(n string, s byte, a int) {
	person.name = n
	person.sex = s
	person.age = a
	fmt.Printf("SetInfoValue %p\n", &person)
	fmt.Println("SetInfoValue", person)
}
func (person *Person) SetInfoPointer(n string, s byte, a int) {
	person.name = n
	person.sex = s
	person.age = a
	fmt.Printf("setInfoPointer %p \n", person)
	fmt.Println("SetInfoPointer", *person)
}
func main() {
	var s Person
	fmt.Printf("main s :%p\n", &s)
	s.SetInfoValue("mike", 'm', 18)
	fmt.Println(s)

	s.SetInfoPointer("yoyo", 'f', 20)
	fmt.Println(s)
}
