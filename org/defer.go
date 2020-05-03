package main

import "fmt"

func main() {
	var i int = 1
	defer fmt.Println("world", i)
	i++
	fmt.Println("hello", i)
}
