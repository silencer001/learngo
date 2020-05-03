package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		for t := range timer.C {
			fmt.Println(t)
		}

		fmt.Println("子协程时间到")
	}()
	time.Sleep(5 * time.Second)

	if !timer.Stop() {
		//fmt.Println(<-timer.C)
	}
	timer.Reset(12 * time.Second)

	for {
	}
}
