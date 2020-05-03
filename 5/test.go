package main

import "fmt"

func main() {
	var a [3]int
	x := 3
	a[x] = 10
	fmt.Println(a[2])
}
