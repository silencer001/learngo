package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "golang", 3: "jude"}

	for key, value := range m {
		fmt.Println(key, ":", value)
	}
	//如何判断一个key是否存在,存在时为true，不存在时为false
	value, ok := m[4]
	if !ok {
		m[4] = "error"
		fmt.Println("m[4] is not exist!")
	} else {
		fmt.Println("value :", value)
	}
	fmt.Println(m)
}
