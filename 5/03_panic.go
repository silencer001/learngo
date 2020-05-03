package main

import "fmt"

func testa() {
	fmt.Println("testa")
}
func testb() {
	fmt.Println("testb")
	panic("at testb panic")
}
func testc() {
	fmt.Println("testc")
}
func main() {
	testa()
	testb()
	testc()
}
