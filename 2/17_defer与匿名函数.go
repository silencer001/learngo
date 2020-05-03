package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Println("a=", a)
		fmt.Println("b=", b)
	}(a, b)
	a = 100
	b = 200
	fmt.Println("外部a=", a)
	fmt.Println("外部b=", b)
}
