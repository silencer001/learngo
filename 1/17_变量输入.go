package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入变量a: ")
	//fmt.Scanf("%d", &a) //取地址
	if i, err := fmt.Scan(&a); err != nil {
		fmt.Printf("error type!\n")
	} else {
		fmt.Printf("input len:%d\n", i)
	}
	fmt.Printf("a = %d\n", a)
}
