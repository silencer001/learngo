package main

import "fmt"

func main() {
	//有多少个方括号就是多少维
	//有多少个方括号就有多少个循环
	var a [3][4]int
	k := 1
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			k++
			a[i][j] = k
		}
	}
	fmt.Println("a = ", a)

	b := [3][4]int{
		{1, 2, 3, 4},
		//右大括号必须在同一行
		{5, 6, 7, 9}} // syntax error: unexpected newline, expecting comma or }
	fmt.Println("b= ", b)
}
