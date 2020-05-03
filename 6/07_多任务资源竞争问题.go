package main

import (
	"fmt"
	"time"
)

//全局变量，创建一个chan
var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second) //sleep时会让出调度
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
	ch <- 666
}
func person2() {
	<-ch
	Printer("world")
}
func main() {
	//runtime.GOMAXPROCS(1)
	// ss := []string{"hhhhhhhhhhhhhhhhhhhhhhh", "LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL", "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBB", "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}
	// for _, s := range ss {
	// 	go Printer(s)
	// }
	go person1()
	go person2()
	for {

	}
}
