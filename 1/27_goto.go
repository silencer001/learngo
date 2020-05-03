package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 5 {
			goto i5 //goto不能跨函数使用
		} else {
			fmt.Println("i=", i)
		}
	}
i5:
	fmt.Println("This is goto :", i)
}
