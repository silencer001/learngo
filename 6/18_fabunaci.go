package main

import (
	"fmt"
	"time"
)

func fab(out chan<- int, end <-chan int) {
	x, y := 1, 1
	for {
		select {
		case out <- x:
			x, y = y, x+y
		case <-end:
			fmt.Println("fab end")
			return
		}
	}
}
func main() {
	ch := make(chan int)
	end := make(chan int)
	go fab(ch, end)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	end <- 666
	time.Sleep(time.Second * 2)
}
