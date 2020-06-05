package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(i interface{}) {
	// if v, ok := i.(int); ok {
	// 	//use of .(type) outside type switch
	// 	//fmt.Printf("input type:%T, value:%d\n", i.(type), v)
	// 	fmt.Printf("input type:%T, value:%d\n", v, v)
	// 	return 1
	// } else {
	// 	fmt.Printf("unkown type,%+v\n", i)
	// 	return 0
	// }
	switch v := i.(type) {
	case int:
		fmt.Printf("i is interger, value:%d\n", v)
	case string:
		fmt.Printf("i is string, value:%s\n", v)
	default:
		fmt.Printf("Unkown type: %+v\n", i)
	}
}

func TestInterface(t *testing.T) {
	//	t.Log(DoSomething(123))
	DoSomething(123)
	DoSomething("hello world")
	DoSomething(123.3)
}
