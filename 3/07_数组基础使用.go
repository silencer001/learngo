package main

import "fmt"

func main() {
	var a [10]int
	const b = 5
	//b:=5
	var bb [b]int //non-constant array bound b
	fmt.Printf("a: %v\n", a)
	fmt.Printf("bb :%v\n", bb)
	//定义一个数据，[10]int和[5]int是两种不同的类型
	fmt.Printf("a type: %T, bb type :%T\n", a, bb)

	//range迭代
	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}
}
