package main

import "fmt"

func main() {
	var a bool
	a = true
	fmt.Println("a=", a)

	var b = false
	fmt.Println("b=", b)
	c := false
	fmt.Println("c=", c)
	fmt.Printf("c=%b", c)
}
