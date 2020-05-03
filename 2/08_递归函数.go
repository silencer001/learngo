package main

import "fmt"

func funca(a int) int {
	a--
	fmt.Println("funca = ", a)
	if a > 0 {
		return funca(a)
	}
	return 0
}
func main() {
	fmt.Println("main func")
	funca(10)
	//fmt.Printf("main a = %d")
}
