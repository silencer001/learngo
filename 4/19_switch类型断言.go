package main

import "fmt"

type Student struct {
	name string
	id   int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student:{name:%s, id:%d}", s.name, s.id)
}
func main() {
	var i []interface{} = make([]interface{}, 3)
	i[0] = 123
	i[1] = "hello"
	i[2] = &Student{"mike", 123}
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("i[%d] type is int, value:%d\n", index, value)
		case string:
			fmt.Printf("i[%d] type is string, value:%s\n", index, value)
		case *Student:
			fmt.Printf("i[%d] type is Student, value:%s\n", index, value)
		default:
			fmt.Printf("error type\n")
		}
	}
}
