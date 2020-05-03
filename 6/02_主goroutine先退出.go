package main

import (
	"fmt"
	"time"
)

func main() {
	//当主协程退出，其他子协程也要跟着退出

	go func() {
		i := 0
		for {
			i++
			fmt.Println("go routine i = ", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
