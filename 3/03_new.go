package main

import "fmt"

func main() {
	var p *int
	fmt.Printf("p = %v\n", p)
	p = new(int)
	*p = 666
	fmt.Printf("p = %v\n", p)
}
