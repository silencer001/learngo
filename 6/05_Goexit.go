package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccc")
	//return //终止此函数
	runtime.Goexit()
	fmt.Println("dddd")
}
func main() {
	go func() {
		defer fmt.Println("test")
		fmt.Println("aaa")

		//调用别的函数
		test()
		fmt.Println("bbb")
	}()
	for {

	}
}
