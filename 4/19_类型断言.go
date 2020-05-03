package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello, go"
	i[2] = &Student{"mike", 666}

	//类型断言
	//第一个返回下标，第二个返回value
	for index, data := range i {
		//第一个返回值，第二个返回判断结果的真假
		if value, ok := data.(int); ok {
			fmt.Printf("x[%d] type is int, value :%d\n", index, value)
		} else if value, ok := data.(string); ok {
			fmt.Printf("x[%d] type is string ,value :%s\n", index, value)
			//必须是指针类型
		} else if value, ok := data.(*Student); ok {
			fmt.Printf("x[%d] type is Student, value:%v\n", index, value)
		} else {
			panic("unkonw type")
		}
	}
}
