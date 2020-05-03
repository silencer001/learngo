package main

//函数也是一种数据类型，可以像整形/浮点型那样赋值，应用于多态
import "fmt"

func add(i int, j int) int {
	return i + j
}
func min(i int, j int) int {
	return i - j
}

type ftesttype func(int, int) int

func main() {
	fmt.Println("add= ", add(1, 2))

	var f ftesttype
	f = add
	fmt.Println("f2 = ", f(10, 20))
	f = min
	fmt.Println("f3 = ", f(10, 5))

	ff := add
	fmt.Println("ff = ", ff(10, 20))
}
