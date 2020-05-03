package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}
	//支持比较，只比较== 或!=，每个元素要完全相同
	fmt.Println("a == b", a == b)
	fmt.Println("a == c", a == c)

	//同类型的数组可以赋值
	var d [5]int
	d = a
	fmt.Println("d = ", d)
}
