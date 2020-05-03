package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)

	fmt.Printf("len(ch) = %d, cap(ch)= %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程：i= ", i)
			ch <- i
			//fmt.Printf("len(ch) = %d, cap(ch)= %d\n", len(ch), cap(ch))
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //读chan，没有内容阻塞
		fmt.Println("主协程： num = ", num)
	}
}
