package main

import "fmt"

func main() {
	var b = 12
	fmt.Println("b=", b)
	var a int = 10
	fmt.Println("a=", a)
	
	var c = 10.0
	//fmt.Println("c=%T",c)
	fmt.Printf("c tyep:%T\n",c)
	d := 11
	fmt.Printf("d : %V\n",d)
}