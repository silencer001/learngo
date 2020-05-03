package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("name: %s, age: %d", s.name, s.age)
}
func main() {
	s := &Student{"mike", 18}
	//空接口实际上是万能类型，可以保存任何类型的值
	fmt.Println("s = ", s)
}
