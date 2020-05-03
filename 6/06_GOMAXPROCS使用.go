package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(4)
	fmt.Println("n=", n)
	for {
		go fmt.Println(1)
		fmt.Print(0)
	}
}
