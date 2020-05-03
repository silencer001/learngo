package main

import "fmt"

func testdel(m map[int]string) {
	delete(m, 1)
}
func main() {
	m := map[int]string{1: "a", 2: "b", 3: "go"}
	fmt.Println("m = ", m)

	//说明map作为形参，传递的是引用
	testdel(m)

	fmt.Println("m = ", m)
}
