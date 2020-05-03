package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		timer := time.NewTimer(3 * time.Second)
		for {
			select {
			case num := <-ch:
				fmt.Println("num=", num)
				timer.Reset(3 * time.Second)
				//time.After将导致严重的资源泄露,正确的做法应该使用NewTimer
			//case <-time.After(3 * time.Second):
			case <-timer.C:
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	<-quit
	fmt.Println("end！")
}
