package main

import "fmt"

func testa() {
	fmt.Println("testa")
}
func testb(x int) {
	defer func() {
		//fmt.Println(recover())
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var array [10]int
	fmt.Println("testb")
	array[x] = 10

}
func testc() {
	fmt.Println("testc")
}

func main() {
	testa()
	testb(11)
	testc()
}
